package restful

import (
	"errors"
	"net/http"
	"web-server/exceptions"
	"web-server/logger"
	"web-server/restful/middlewares"
	"web-server/restful/request"
	"web-server/restful/response"
	"web-server/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userController struct {
	userService service.UserServiceWithContext
}

func NewUserController(group *gin.RouterGroup, authChecker middlewares.AuthChecker, userService service.UserServiceWithContext) {
	controller := &userController{
		userService: userService,
	}
	v1 := group.Group("/v1")
	v1.GET("/user/:id", authChecker.Check, controller.GetUserByIdHandler)
	v1.GET("/users", authChecker.Check, controller.GetUsersHandler)
	v1.PUT("/user/:id", authChecker.Check, controller.UpdateUserHandler)
	v1.DELETE("/user/:id", authChecker.Check, controller.DeleteUserHandler)
	v1.POST("/signup", controller.CreateUserHandler)
}

func (ctrl *userController) GetUserByIdHandler(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, exceptions.UUIDParseFailed.SetMessage(ctx.Param("id")))
		return
	}
	user, err := ctrl.userService(ctx).GetUserById(id)
	if err != nil {
		if errors.Is(err, exceptions.UserNotExists) {
			ctx.JSON(http.StatusNotFound, err)
		} else {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		return
	}
	ctx.JSON(http.StatusOK, response.NewGetUserResponse(user))
}

func (ctrl *userController) GetUsersHandler(ctx *gin.Context) {
	page := request.NewPageRequest(ctx.Query("page"), ctx.Query("pageSize"))
	users, totalCount, err := ctrl.userService(ctx).ListUsers(page.Page, page.PageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, response.NewGetUsersResponse(users, response.NewPagination(page.Page, page.PageSize, totalCount)))
}

func (ctrl *userController) CreateUserHandler(ctx *gin.Context) {
	var createUserRequest request.CreateUserRequest
	err := ctx.ShouldBindJSON(&createUserRequest)
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	user, err := createUserRequest.ToUser()
	if err != nil {
		logger.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	if user, err = ctrl.userService(ctx).Create(user); err != nil {
		if errors.Is(err, exceptions.UserAlreadyExist) {
			ctx.JSON(http.StatusCreated, err)
		} else {
			ctx.JSON(http.StatusInternalServerError, err)
		}
	} else {
		ctx.JSON(http.StatusOK, response.NewCreateUserResponse(user))
	}
}

func (ctrl *userController) LoginHandler(ctx *gin.Context) {
	var loginRequest request.EmailLoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, exceptions.UUIDParseFailed)
	}

}

func (ctrl *userController) UpdateUserHandler(ctx *gin.Context) {
	var updateRequest request.UpdateUserRequest
	_, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, exceptions.UUIDParseFailed)
		return
	}
	if err = ctx.ShouldBindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, exceptions.UUIDParseFailed)
	}
	user, err := updateRequest.ToUser()
	if err = ctrl.userService(ctx).Update(user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
	} else {
		ctx.Status(http.StatusOK)
	}
}

func (ctrl *userController) DeleteUserHandler(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, exceptions.UUIDParseFailed)
	}
	err = ctrl.userService(ctx).DeleteUserById(id)
	if err != nil {
		if errors.Is(err, exceptions.UserNotExists) {
			ctx.JSON(http.StatusNotFound, err)
		} else {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		return
	}
	ctx.Status(http.StatusOK)
}
