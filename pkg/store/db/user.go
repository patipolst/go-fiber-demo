package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/patipolst/go-fiber-demo/pkg/user"
)

type UserModel struct {
	gorm.Model
	user.User
}

func NewUserModel(u user.User) UserModel {
	return UserModel{
		User: u,
	}
}

func (m UserModel) TableName() string {
	return "users"
}

func (m UserModel) ToUser() user.User {
	return user.User{
		Name: m.User.Name,
		Age:  m.User.Age,
	}
}

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

	s.db.AutoMigrate(&UserModel{})

	return s, nil
}

func (s *Store) GetAllUsers() []user.User {
	var models []UserModel
	var users []user.User
	s.db.Find(&models)
	for _, m := range models {
		users = append(users, m.ToUser())
	}
	return users
}

func (s *Store) GetUser(id int) user.User {
	var m UserModel
	s.db.First(&m, id)
	return m.ToUser()
}

func (s *Store) CreateUser(user user.User) user.User {
	m := NewUserModel(user)
	s.db.Create(&m)
	return user
}

func (s *Store) DeleteUser(id int) {
	var m UserModel
	s.db.First(&m, id)
	s.db.Delete(&m)
}
