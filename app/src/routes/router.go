package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter Router creation
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		var muxRoute *mux.Route
		muxRoute = router.Methods(route.Method)
		muxRoute.
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
