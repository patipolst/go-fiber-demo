package stub

import (
	"github.com/jinzhu/gorm"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

type Store struct {
	db *gorm.DB
}

func NewStore() (*Storage, error) {
	s.db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) GetAllUsers() []user.User {
	var users []user.User
	db.Find(&value)
	return users
}

func (s *Store) GetUser(id int) user.User {
	var user user.User
	db.First(&user, id)
	return user
}

func (s *Store) CreateUser(user user.User) user.User {
	db.Create(&User)
	return user
}

func (s *Store) DeleteUser(id int) {
	var user user.User
	db.First(&user, id)
	db.Delete(&user)
}
