package main

import (
	"Api1/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/user/delete/{id}", handler.DeleteUser)
	http.ListenAndServe(":8030", nil)
}
