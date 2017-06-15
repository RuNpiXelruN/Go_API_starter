package controller

import (
	"fmt"
	"net/http"
)

type about struct{}

func (a about) registerRoutes() {
	http.HandleFunc("/about", a.handleAbout)
}

func (a about) handleAbout(w http.ResponseWriter, req *http.Request) {
	fmt.Println("about handler hit :)")
}
