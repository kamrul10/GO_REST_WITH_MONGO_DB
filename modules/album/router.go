
package album

import (
	"log"
	"net/http"
	"musicstore/libs/logger"
	"github.com/gorilla/mux"
	"gopkg.in/matryer/respond.v1"
)
const version = "1.0"
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
	//response for different status
	opts := &respond.Options{
		Before: func(w http.ResponseWriter, r *http.Request, status int, data interface{}) (int, interface{}) {
			w.Header().Set("X-API-Version", version)
			dataenvelope := map[string]interface{}{"code": status}
			if err, ok := data.(error); ok {
				dataenvelope["success"] = false
				dataenvelope["error"] = err.Error()
				
			} else {
				dataenvelope["success"] = true
				dataenvelope["data"] = data
				
			}
			return status, dataenvelope
 
		},
		After: func(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
			log.Println("<-", status, data)
		},
	}
	router := mux.NewRouter()
	for _,route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(opts.Handler(handler))
	}
	return router
}
