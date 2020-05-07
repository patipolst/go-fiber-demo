package user

import "fmt"

// Service provides user operations.
type Service interface {
	GetAllUsers() []User
	GetUser(id int) User
	CreateUser(user User) User
	DeleteUser(id int)
}

// Store provides access to repository.
type Store interface {
	GetAllUsers() []User
	GetUser(id int) User
	CreateUser(user User) User
	DeleteUser(id int)
}

type service struct {
	store Store
}

// NewService creates an adding service with the necessary dependencies
func NewService(store Store) Service {
	return &service{store}
}

func (s *service) GetAllUsers() []User {
	// any validation can be done here
	fmt.Println("SERVICE: GET all users")
	return s.store.GetAllUsers()
}

func (s *service) GetUser(id int) User {
	return s.store.GetUser(id)
}

func (s *service) CreateUser(user User) User {
	return s.store.CreateUser(user)
}

func (s *service) DeleteUser(id int) {
	s.store.DeleteUser(id)
}
