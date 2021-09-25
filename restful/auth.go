package restful

import (
	"errors"
	"net/http"
	"web-server/exceptions"
	"web-server/restful/request"
	"web-server/restful/response"
	"web-server/service"

	"github.com/gin-gonic/gin"
)

type authController struct {
	authService service.AuthServiceWithContext
}

func NewAuthController(group *gin.RouterGroup, authService service.AuthServiceWithContext) *authController {
	controller := &authController{
		authService: authService,
	}
	v1 := group.Group("/v1")
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/refresh_token", controller.RefreshTokenHandler)
	return controller
}

func (ctrl *authController) LoginHandler(ctx *gin.Context) {
	var loginRequest request.EmailLoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	token, err := ctrl.authService(ctx.Request.Context()).AuthenticateByEmail(loginRequest.Email, loginRequest.Password)
	if err != nil {
		if errors.Is(err, exceptions.UserNotExists) {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	ctx.JSON(http.StatusOK, response.NewLoginResponse(token))
}

func (ctrl *authController) RefreshTokenHandler(ctx *gin.Context) {
	var refreshTokenRequest request.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&refreshTokenRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	token, err := ctrl.authService(ctx.Request.Context()).RefreshToken(refreshTokenRequest.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}
	ctx.JSON(http.StatusOK, response.NewRefreshTokenResponse(token))
}
