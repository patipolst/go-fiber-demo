package main

import (
	"github.com/patipolst/go-fiber-demo/pkg/http/rest"
	"github.com/patipolst/go-fiber-demo/pkg/store/db"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

func main() {
	// userStore := stub.NewStore()
	userStore, _ := db.NewStore()
	userService := user.NewService(userStore)
	app := rest.Handler(userService)
	app.Listen(3000)
}
