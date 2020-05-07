package user

import "github.com/jinzhu/gorm"

// User represents user model
type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}
