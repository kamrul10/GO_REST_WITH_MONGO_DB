
package album

import (
	"net/http"
	"musicstore/libs/logger"
	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"GetAlbums",
		"GET",
		"/albums",
		controller.GetAlbums,

	},Route{
		"GetAlbum",
		"GET",
		"/albums/{id}",
		controller.GetAlbum,

	},
	Route{
		"AddAlbum",
		"POST",
		"/albums",
		controller.AddAlbum,
	},
	Route{
		"UpdateAlbum",
		"PUT",
		"/albums",
		controller.UpdateAlbum,
	},
	Route{
		"DeleteAlbum",
		"DELETE",
		"/albums/{id}",
		controller.DeleteAlbum,
	},
}
func NewRouter() *mux.Router {

	router := mux.NewRouter()
	for _,route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
