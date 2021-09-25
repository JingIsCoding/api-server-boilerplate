package response

import "web-server/model"

type CreateUserResponse struct {
	ID string `json:"id"`
}

type UpdateUsersResponse struct{}

type GetUserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type GetUsersResponse struct {
	Pagination Pagination   `json:"pagination"`
	Users      []model.User `json:"users"`
}

func NewCreateUserResponse(user model.User) CreateUserResponse {
	return CreateUserResponse{
		ID: user.ID.String(),
	}
}

func NewGetUserResponse(user model.User) GetUserResponse {
	return GetUserResponse{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
}

func NewGetUsersResponse(users []model.User, pagination Pagination) GetUsersResponse {
	return GetUsersResponse{
		Users:      users,
		Pagination: pagination,
	}
}
