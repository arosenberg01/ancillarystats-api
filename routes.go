package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler http.Handler
}
type Routes []Route
type appHandler func(http.ResponseWriter, *http.Request) (int, error)

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status, err := fn(w, r); err != nil {
		log.Print(err)

		switch status {
			case http.StatusNotFound:
				http.NotFound(w, r)
			case http.StatusInternalServerError:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		wrappedHandler := Logger(route.Handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(wrappedHandler)
	}

	return router
}

var routes = Routes{
	Route{
		"Player",
		"GET",
		"/player/{player_id}",
		appHandler(PlayerHandler),
	},
	Route{
		"Leaders",
		"GET",
		"/leaders/{category}",
		appHandler(LeadersHandler),
	},
}
