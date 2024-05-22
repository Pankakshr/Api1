package main

import (
	"Api1/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/user/create", handler.CreateUser)
	http.ListenAndServe(":8030", nil)
}
