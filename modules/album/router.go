
package album

import (
	"net/http"
	"musicstore/libs/logger"
	"github.com/gorilla/mux"
	"gopkg.in/matryer/respond.v1"
	"musicstore/types"
	"musicstore/middlewares"
)
const version = "1.0"
var controller = &Controller{Repository: Repository{}}
type Route = types.Route
type AlbumRoutes []Route

var routes = AlbumRoutes{
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


func AlbumRouter(router *mux.Router){
	opts := &respond.Options{
		Before: func(w http.ResponseWriter, r *http.Request, status int, data interface{}) (int, interface{}) {
			w.Header().Set("X-API-Version", version)
			dataenvelope := map[string]interface{}{}
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
			// log.Println("<-", status, data)
		},
	}

	for _,route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(middlewares.JwtMiddleware(opts.Handler(handler)))
	}


}


