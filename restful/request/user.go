package request

import (
	"web-server/exceptions"
	"web-server/model"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	Phone           string         `json:"phone"`
	Email           string         `json:"email"`
	Password        model.Password `json:"password"`
	ConfirmPassword model.Password `json:"confirm_password"`
}

type UpdateUserRequest struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (createUserRequest CreateUserRequest) Validate() error {
	if len(createUserRequest.Email) == 0 {
		return exceptions.UserCreateFaled.SetMessage("email is empty")
	}
	if len(createUserRequest.Password) == 0 {
		return exceptions.UserCreateFaled.SetMessage("password is empty")
	}
	if createUserRequest.Password != createUserRequest.ConfirmPassword {
		return exceptions.UserCreateFaled.SetMessage("password not match")
	}
	return nil
}

func (createUserRequest CreateUserRequest) ToUser() (model.User, error) {
	if err := createUserRequest.Validate(); err != nil {
		return model.User{}, err
	}
	return model.User{
		FirstName:         createUserRequest.FirstName,
		LastName:          createUserRequest.LastName,
		Phone:             createUserRequest.Phone,
		Email:             createUserRequest.Email,
		EncryptedPassword: createUserRequest.Password.Encrypt(),
	}, nil
}

func (updateRequest UpdateUserRequest) Validate() error {
	if len(updateRequest.Email) == 0 {
		return exceptions.UserCreateFaled.SetMessage("email is empty")
	}
	return nil
}

func (updateRequest UpdateUserRequest) ToUser() (model.User, error) {
	if err := updateRequest.Validate(); err != nil {
		return model.User{}, err
	}
	return model.User{
		Base: model.Base{
			ID: uuid.MustParse(updateRequest.ID),
		},
		FirstName: updateRequest.FirstName,
		LastName:  updateRequest.LastName,
		Phone:     updateRequest.Phone,
		Email:     updateRequest.Email,
	}, nil
}
