package service

type UserStore interface {
	GetUsers() []User
	GetUser(id int) User
	CreateUser(user User) User
	DeleteUser(id int)
}

type StubUserStore struct {
	users map[int]User
}

func NewDummyUserStore() *StubUserStore {
	return &StubUserStore{
		map[int]User{
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

func (s *StubUserStore) GetUsers() []User {
	var users []User
	for _, value := range s.users {
		users = append(users, value)
	}
	return users
}

func (s *StubUserStore) GetUser(id int) User {
	user := s.users[id]
	return user
}

func (s *StubUserStore) CreateUser(user User) User {
	s.users[10] = user // fixme
	return user
}

func (s *StubUserStore) DeleteUser(id int) {
	delete(s.users, id)
}
