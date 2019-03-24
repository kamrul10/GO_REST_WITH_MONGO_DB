package routes
import (
	"github.com/gorilla/mux"
	"musicstore/modules/album"
	
)
func NewRouter() *mux.Router {
	//response for different status
	router := mux.NewRouter()
	album.AlbumRouter(router)
	return router
}