package controller

import (
	"fmt"
	"net/http"

	"../model"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, req *http.Request) {
	user := model.GetUser()
	fmt.Println("user :)", user.FirstName)
}
