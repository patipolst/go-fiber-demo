package service

// User represents user model
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserService struct {
	store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{store}
}

func (s *UserService) GetUsers() []User {
	return s.store.GetUsers()
}

func (s *UserService) GetUser(id int) User {
	return s.store.GetUser(id)
}

func (s *UserService) CreateUser(user User) User {
	return s.store.CreateUser(user)
}

func (s *UserService) DeleteUser(id int) {
	s.store.DeleteUser(id)
}
