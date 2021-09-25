package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"web-server/config"
	"web-server/exceptions"
	"web-server/model"
	"web-server/repo"

	"github.com/golang-jwt/jwt"
)

const tokenExpirationDuration = 30 //minutes

type AuthServiceWithContext func(ctx context.Context) AuthService

type AuthService interface {
	AuthenticateByEmail(email string, password string) (token string, err error)
	Validate(token string, claims jwt.MapClaims) (*jwt.Token, error)
	RefreshToken(token string) (newToken string, err error)
}

type authServiceImpl struct {
	ctx      context.Context
	config   *config.Config
	userRepo repo.UserRepoWithContext
}

func (authService *authServiceImpl) AuthenticateByEmail(email string, password string) (string, error) {
	user, err := authService.userRepo(authService.ctx).GetByEmail(email)
	hmacSecret := authService.config.ServerConfig.HMACSecret
	domain := authService.config.ServerConfig.Domain
	if err != nil {
		return "", exceptions.UserNotExists
	}
	if !user.EncryptedPassword.Compare(model.Password(password)) {
		return "", exceptions.AuthFailed.SetMessage("email or password not match")
	}
	claims := jwt.MapClaims{
		"aud": domain,
		"iss": domain,
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Minute * tokenExpirationDuration).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
	}
	return newToken(hmacSecret, claims)
}

func (authService *authServiceImpl) Validate(tokenString string, checkClaims jwt.MapClaims) (*jwt.Token, error) {
	token, err := parseToken(tokenString, authService.config.ServerConfig.HMACSecret)
	if err != nil {
		return nil, exceptions.InvalidToken.Wrap(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		exp := claims["exp"].(float64)
		if time.Now().Unix() > int64(exp) {
			return nil, exceptions.TokenExpires
		}
		for claimKey, claimValue := range checkClaims {
			if value, ok := claims[claimKey]; !ok || value != claimValue {
				return nil, exceptions.AuthFailed.SetMessage("failed to auth")
			}
		}
		return token, nil
	}
	return nil, exceptions.InvalidToken
}

func (authService *authServiceImpl) RefreshToken(tokenString string) (string, error) {
	var token *jwt.Token
	var err error
	hmacSecret := authService.config.ServerConfig.HMACSecret
	if token, err = authService.Validate(tokenString, jwt.MapClaims{}); err != nil {
		if !errors.Is(err, exceptions.TokenExpires) {
			return "", err
		}
	}
	return newToken(hmacSecret, token.Claims)
}

func newToken(hmacSecret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", exceptions.InvalidToken.SetMessage(err.Error())
	}
	return tokenString, nil
}

func parseToken(tokenString string, hmacSecret string) (*jwt.Token, error) {
	parser := &jwt.Parser{
		SkipClaimsValidation: true,
	}
	token, err := parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSecret), nil
	})
	if err != nil {
		return nil, exceptions.AuthFailed.SetMessage("failed to parse token").Wrap(err)
	}
	if token.Valid {
		return token, nil
	}
	return nil, exceptions.AuthFailed.SetMessage("failed to validate token")
}

func NewAuthServiceWithContext(userRepo repo.UserRepoWithContext) AuthServiceWithContext {
	config := config.Get()
	return func(ctx context.Context) AuthService {
		return &authServiceImpl{
			config:   config,
			userRepo: userRepo,
			ctx:      ctx,
		}
	}
}
