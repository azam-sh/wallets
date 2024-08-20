package router

import (
	"wallets/internal/transport/http/handlers"
	"wallets/internal/transport/http/middleware"

	"github.com/gorilla/mux"
)

func InitRouter(handlers *handlers.Handler, mw middleware.MiddlewareInterface) *mux.Router {
	router := mux.NewRouter()

	return router
}
