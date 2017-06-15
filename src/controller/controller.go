package controller

var (
	homeController  home
	aboutController about
)

// Startup function to register each controller's routes
func Startup() {
	homeController.registerRoutes()
	aboutController.registerRoutes()

	// http.Handle("/img/", http.FileServer(http.Dir("public")))
	// http.Handle("/css/", http.FileServer(http.Dir("public")))
}
