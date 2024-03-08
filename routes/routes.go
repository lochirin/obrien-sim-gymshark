package routes

import (
	"example/gymshark/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func PackageRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/packages", controller.ComputePackages).Methods(http.MethodPost)

	return router
}
