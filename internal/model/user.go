package model

import "github.com/aurelius15/lf/internal/utils"

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type User struct {
	ID                 string
	Name               string
	Balance            int
	VerificationStatus bool
}

func CreateUser(name string) *User {
	return &User{
		ID:                 utils.GenerateUUID(),
		Name:               name,
		Balance:            1000,
		VerificationStatus: false,
	}
}
