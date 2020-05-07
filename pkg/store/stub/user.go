package stub

import "github.com/patipolst/go-fiber-demo/pkg/user"

type Store struct {
	users map[int]user.User
}

func NewStore() *Store {
	return &Store{
		map[int]user.User{
			1: {
				Name: "A",
				Age:  10,
			},
			2: {
				Name: "B",
				Age:  20,
			},
		},
	}
}

func (s *Store) GetAllUsers() []user.User {
	var users []user.User
	for _, value := range s.users {
		users = append(users, value)
	}
	return users
}

func (s *Store) GetUser(id int) user.User {
	user := s.users[id]
	return user
}

func (s *Store) CreateUser(user user.User) user.User {
	s.users[10] = user // fixme
	return user
}

func (s *Store) DeleteUser(id int) {
	delete(s.users, id)
}
