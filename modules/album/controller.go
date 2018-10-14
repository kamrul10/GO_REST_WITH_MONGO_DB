package album

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"gopkg.in/matryer/respond.v1"
	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	Repository Repository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusOK, "welcome to the era of golang")
}

//GET albums
func (c *Controller) GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums := c.Repository.GetAlbums() // list of all albums
	log.Println(albums)
	respond.With(w, r, http.StatusOK, albums)
}

//Get Album
func (c *Controller) GetAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]                                   // param id
	album := c.Repository.GetAlbum(id)
	log.Println(album)
	respond.With(w, r, http.StatusOK, album)
}
// AddAlbum POST /
func (c *Controller) AddAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error AddAlbum", err)
		respond.With(w, r, http.StatusInternalServerError, err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddAlbum", err)
	}
	if err := json.Unmarshal(body, &album); err != nil { // unmarshall body contents as a type Candidate
		respond.With(w, r, 422, err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddAlbum unmarshalling data", err)
			respond.With(w, r, http.StatusInternalServerError, err)
		}
	}
	success := c.Repository.AddAlbum(album) // adds the album to the DB
	if !success {
		respond.With(w, r, http.StatusInternalServerError, err)
	}
	respond.With(w, r, http.StatusCreated, album)
}

// UpdateAlbum PUT /
func (c *Controller) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error UpdateAlbum", err)
		respond.With(w, r, http.StatusInternalServerError, err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddaUpdateAlbumlbum", err)
	}
	if err := json.Unmarshal(body, &album); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		respond.With(w, r, 422, err) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateAlbum unmarshalling data", err)
			respond.With(w, r, http.StatusBadRequest, err)
		}
	}
	success := c.Repository.UpdateAlbum(album) // updates the album in the DB
	if !success {
		respond.With(w, r, http.StatusInternalServerError, err)
	}
	respond.With(w, r, http.StatusOK, album)
}

// DeleteAlbum DELETE /
func (c *Controller) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]                                    // param id
	if err := c.Repository.DeleteAlbum(id); err != "" { // delete a album by id
		if strings.Contains(err, "404") {
			respond.With(w, r, http.StatusNotFound, err)
			
		} else if strings.Contains(err, "500") {
			respond.With(w, r, http.StatusInternalServerError, err)
		}
	}
	respond.With(w, r, http.StatusOK, "")
}
