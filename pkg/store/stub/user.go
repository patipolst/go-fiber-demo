package stub

import "github.com/patipolst/go-fiber-demo/pkg/user"

type UserStore struct {
	users map[int]user.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		make(map[int]user.User),
	}
}

func (s *UserStore) GetAllUsers() []user.User {
	var users []user.User
	for _, value := range s.users {
		users = append(users, value)
	}
	return users
}

func (s *UserStore) GetUser(id int) user.User {
	user := s.users[id]
	return user
}

func (s *UserStore) CreateUser(user user.User) user.User {
	id := len(s.users) + 1
	s.users[id] = user
	return user
}

func (s *UserStore) DeleteUser(id int) {
	delete(s.users, id)
}
