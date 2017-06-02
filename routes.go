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

func NewRouter(env *Env) *mux.Router {
	var routes = Routes{
		Route{
			"NbaPlayer",
			"GET",
			"/players/{player_id}",
			appHandler(env.NbaPlayerHandler),
		},
		Route{
			"NbaLeaders",
			"GET",
			"/leaders/{category}",
			appHandler(env.NbaLeadersHandler),
		},
		Route{
			"NbaCategories",
			"GET",
			"/categories",
			appHandler(env.NbaCategoriesHandler),
		},
		Route{
			"NbaTeams",
			"GET",
			"/teams",
			appHandler(env.NbaTeamsHandler),
		},
		Route{
			"NbaRosters",
			"GET",
			"/rosters/{team_id}",
			appHandler(env.NbaRosterHandler),
		},
		Route{
			"NbaGames",
			"GET",
			"/games/{player_id}",
			appHandler(env.NbaGamesHandler),
		},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		wrappedHandler := Logger(route.Handler, route.Name)


		router.
			Methods(route.Method).
			Name(route.Name).
			Handler(wrappedHandler).
			PathPrefix("/v1/nba/").
			Path(route.Pattern)
	}

	return router
}


