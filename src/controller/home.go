package controller

import (
	"net/http"

	"../model"
	"../utils"
)

type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, req *http.Request) {
	user := model.GetUser()
	utils.JSON(w, user)
}
