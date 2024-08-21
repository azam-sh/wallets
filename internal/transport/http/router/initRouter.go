package router

import (
	"wallets/internal/transport/http/handlers"
	"wallets/internal/transport/http/middleware"

	"github.com/gorilla/mux"
)

func InitRouter(handlers *handlers.Handler, mw middleware.Middleware) *mux.Router {
	router := mux.NewRouter()
	router.Use(mw.LogMiddleware)
	router.HandleFunc("/ping", handlers.Ping).Methods("GET")

	privateRouter := router.NewRoute().Subrouter()
	privateRouter.Use(mw.Authenticate)
	privateRouter.HandleFunc("/check-acc", handlers.CheckAccount).Methods("POST")
	privateRouter.HandleFunc("/top-up-balance", handlers.TopUpBalance).Methods("POST")
	privateRouter.HandleFunc("/trns-history", handlers.GetMonthlyTrns).Methods("POST")
	privateRouter.HandleFunc("/balance", handlers.GetBalance).Methods("POST")

	return router
}
