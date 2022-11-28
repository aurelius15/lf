package storage

import (
	"fmt"
	"sync"

	"github.com/aurelius15/lf/internal/model"
)

type InMemory struct {
	mu sync.RWMutex
	v  map[string]*model.User
}

type ErrNotFound struct {
	userId string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("user %s not found \n", e.userId)
}

var storage = InMemory{
	v: make(map[string]*model.User),
}

func Instance() *InMemory {
	return &storage
}

func (s *InMemory) SaveUser(user *model.User) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.v[user.ID] = user
}

func (s *InMemory) SetAsVerified(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if user, ok := s.v[id]; ok {
		user.VerificationStatus = true
		s.v[id] = user
		return nil
	}

	return &ErrNotFound{userId: id}
}

func (s *InMemory) FreezeAmount(userId string, amount int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.v[userId]
	if !ok {
		return &ErrNotFound{userId: userId}
	}

	if err := user.Freeze(amount); err != nil {
		return err
	}

	s.v[userId] = user

	return nil
}

func (s *InMemory) TransferMoney(receiverId, senderId string, amount int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	receiver, ok := s.v[receiverId]
	if !ok {
		return &ErrNotFound{userId: receiverId}
	}

	sender, ok := s.v[senderId]
	if !ok {
		return &ErrNotFound{userId: senderId}
	}

	err := sender.ApplyDeduction(amount)
	if err != nil {
		return err
	}

	receiver.ApplyIncome(amount)

	s.v[receiverId] = receiver
	s.v[senderId] = sender

	return nil
}

func (s *InMemory) GetUser(id string) (*model.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if user, ok := s.v[id]; ok {
		return user, nil
	}

	return nil, &ErrNotFound{userId: id}
}

func (s *InMemory) GetAllUsers() []*model.User {
	users := make([]*model.User, 0, len(s.v))
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.v {
		users = append(users, user)
	}

	return users
}
