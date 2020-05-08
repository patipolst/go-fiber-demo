package main

import (
	"github.com/patipolst/go-fiber-demo/pkg/http/rest"
	"github.com/patipolst/go-fiber-demo/pkg/store/db"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

func main() {
	// userStore := stub.NewStore()
	userStore, _ := db.NewUserStore()
	userService := user.NewService(userStore)
	userHandler := rest.NewUserHandler(userService)
	app := rest.New(userHandler)
	app.Listen(3000)
}
