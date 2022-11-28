package storage

import (
	"sync"

	"github.com/aurelius15/lf/internal/model"
	"gopkg.in/errgo.v2/errors"
)

var storage = sync.Map{}

func SaveUser(user *model.User) (bool, error) {
	storage.Store(user.ID, user)
	return true, nil
}

func SetAsVerified(id string) (bool, error) {
	if user, ok := storage.Load(id); ok {
		user.(*model.User).VerificationStatus = true
		storage.Store(id, user)
		return true, nil
	}

	return false, nil
}

func GetUser(id string) (*model.User, error) {
	if user, ok := storage.Load(id); ok {
		return user.(*model.User), nil
	}

	return nil, errors.New("user not found")
}

func GetAllUsers() ([]*model.User, error) {
	var users []*model.User

	storage.Range(func(key, value interface{}) bool {
		users = append(users, value.(*model.User))
		return true
	})

	return users, nil
}
