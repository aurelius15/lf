package model

import (
	"fmt"

	"github.com/aurelius15/lf/internal/utils"
)

var ErrInsufficientBalance = fmt.Errorf("insufficient balance")

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
}

type User struct {
	ID                 string
	Name               string
	Balance            int
	freeze             int
	VerificationStatus bool
}

func CreateUser(name string) *User {
	return &User{
		ID:                 utils.GenerateUUID(),
		Name:               name,
		Balance:            1000,
		freeze:             0,
		VerificationStatus: false,
	}
}

func (u *User) Freeze(amount int) error {
	if (u.Balance - u.freeze) < amount {
		return ErrInsufficientBalance
	}

	u.freeze += amount
	return nil
}

func (u *User) ApplyDeduction(amount int) error {
	if u.Balance < amount {
		return ErrInsufficientBalance
	}

	u.Balance -= amount
	u.freeze -= amount

	return nil
}

func (u *User) ApplyIncome(amount int) {
	u.Balance += amount
}
