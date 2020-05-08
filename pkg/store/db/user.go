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

type UserStore struct {
	db *gorm.DB
}

func NewUserStore() (*UserStore, error) {
	var err error
	s := new(UserStore)

	s.db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}

	s.db.AutoMigrate(&UserModel{})

	return s, nil
}

func (s *UserStore) GetAllUsers() []user.User {
	var models []UserModel
	var users []user.User
	s.db.Find(&models)
	for _, m := range models {
		users = append(users, m.ToUser())
	}
	return users
}

func (s *UserStore) GetUser(id int) user.User {
	var m UserModel
	s.db.First(&m, id)
	return m.ToUser()
}

func (s *UserStore) CreateUser(user user.User) user.User {
	m := NewUserModel(user)
	s.db.Create(&m)
	return user
}

func (s *UserStore) DeleteUser(id int) {
	var m UserModel
	s.db.First(&m, id)
	s.db.Delete(&m)
}
