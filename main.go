package main

import (
	"beginner-server/app"
	"beginner-server/app/actions"
	"fmt"
	"net/http"
)

const (
	HOST = "localhost"
	PORT = 12345
)

func addr() string {
	return fmt.Sprintf("%s:%d", HOST, PORT)
}

func main() {
	actions.Register()
	fmt.Printf("Listening Server At: %s \n", addr())
	http.ListenAndServe(addr(), nil)
	defer app.CloseDb()
}
