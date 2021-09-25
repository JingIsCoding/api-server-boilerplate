package middlewares

import (
	"net/http"
	"strings"
	"web-server/exceptions"
	"web-server/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthChecker interface {
	Check(*gin.Context)
}

type authCheckImpl struct {
	authService service.AuthServiceWithContext
}

func (checker *authCheckImpl) Check(ctx *gin.Context) {
	context := ctx.Request.Context()
	authorization := ctx.GetHeader("Authorization")
	parts := strings.Split(authorization, " ")
	if len(parts) < 2 {
		ctx.JSON(http.StatusUnauthorized, exceptions.AuthFailed)
		ctx.Abort()
		return
	}
	_, err := checker.authService(context).Validate(parts[1], jwt.MapClaims{})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		ctx.Abort()
		return
	}
}

func NewAuthChecker(authService service.AuthServiceWithContext) AuthChecker {
	return &authCheckImpl{
		authService: authService,
	}
}
