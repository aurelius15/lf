package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user := CreateUser("test")

	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Balance, 1000)
}

func TestUser_Freeze(t *testing.T) {
	u := CreateUser("test")
	err := u.Freeze(500)

	assert.Equal(t, err, nil)
	assert.Equal(t, u.Balance, 1000)
	assert.Equal(t, u.freeze, 500)

	err = u.Freeze(600)
	assert.Error(t, err, ErrInsufficientBalance)
	assert.Equal(t, u.Balance, 1000)
	assert.Equal(t, u.freeze, 500)
}

func TestUser_ApplyDeduction(t *testing.T) {
	u := CreateUser("test")
	fErr := u.Freeze(500)
	dErr := u.ApplyDeduction(500)

	assert.Nil(t, fErr)
	assert.Nil(t, dErr)
	assert.Equal(t, u.Balance, 500)
	assert.Equal(t, u.freeze, 0)

	err := u.ApplyDeduction(600)
	assert.Error(t, err, ErrInsufficientBalance)
}

func TestUser_ApplyIncome(t *testing.T) {
	u := CreateUser("test")
	u.ApplyIncome(500)

	assert.Equal(t, u.Balance, 1500)
	assert.Equal(t, u.freeze, 0)
}
