package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/patipolst/go-fiber-demo/pkg/user"
)

type Store struct {
	db *gorm.DB
}

func NewStore() (*Store, error) {
	var err error
	s := new(Store)

	s.db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}

	s.db.AutoMigrate(&user.User{})

	return s, nil
}

func (s *Store) GetAllUsers() []user.User {
	var users []user.User
	s.db.Find(&users)
	return users
}

func (s *Store) GetUser(id int) user.User {
	var user user.User
	s.db.First(&user, id)
	return user
}

func (s *Store) CreateUser(user user.User) user.User {
	s.db.Create(&user)
	return user
}

func (s *Store) DeleteUser(id int) {
	var user user.User
	s.db.First(&user, id)
	s.db.Delete(&user)
}
