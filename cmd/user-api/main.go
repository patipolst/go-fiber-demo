package main

import (
	"github.com/patipolst/go-fiber-demo/pkg/http/rest"
	"github.com/patipolst/go-fiber-demo/pkg/store/stub"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

func main() {
	userStore := stub.NewUserStore()
	// userStore, _ := db.NewUserStore()
	userService := user.NewService(userStore)
	userHandler := rest.NewUserHandler(userService)
	app := rest.New(userHandler)
	app.Listen(3000)
}
