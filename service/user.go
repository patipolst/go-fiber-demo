package service

import "github.com/gofiber/fiber"

// GetUsers returns all users
func GetUsers(c *fiber.Ctx) {
	c.Send("All Users")
}

// GetUser returns a single user
func GetUser(c *fiber.Ctx) {
	c.Send("Single User")
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) {
	c.Send("New User")
}

// DeleteUser deletes an existing user
func DeleteUser(c *fiber.Ctx) {
	c.Send("Delete User")
}
