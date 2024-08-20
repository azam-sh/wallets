package router

import (
	"wallets/internal/transport/http/handlers"
	"wallets/internal/transport/http/middleware"

	"github.com/gorilla/mux"
)

func InitRouter(handlers *handlers.Handler, mw middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handlers.Ping).Methods("GET")

	privateRouter := router.NewRoute().Subrouter()
	privateRouter.Use(mw.Authenticate)

	return router
}
