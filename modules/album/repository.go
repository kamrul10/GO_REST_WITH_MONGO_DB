package album

import (
	"fmt"
	"os"
  "musicstore/libs/db"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct {
}


// DBNAME the name of the DB instance
var DBNAME = os.Getenv("MONGO_DATABASE")

// DOCNAME the name of the document
var DOCNAME = "albums"

// GetAlbums returns the list of Albums
func (r *Repository) GetAlbums() Albums {
	session := db.Session()
	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)
	results := Albums{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

//get Album
func (r *Repository) GetAlbum(id string) Album {

	session := db.Session()
	defer session.Close()
	//Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return Album{}
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	c := session.DB(DBNAME).C(DOCNAME)
	result := Albums{}
	if err := c.FindId(oid).All(&result); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return result[0]
}

// AddAlbum inserts an Album in the DB
func (r *Repository) AddAlbum(album Album) bool {
	session := db.Session()
	defer session.Close()
	album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(album)
	return true
}

// UpdateAlbum updates an Album in the DB (not used for now)
func (r *Repository) UpdateAlbum(album Album) bool {
	session := db.Session()
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(album.ID, album)
	return true
}

// DeleteAlbum deletes an Album (not used for now)
func (r *Repository) DeleteAlbum(id string) string {
	session := db.Session()
	defer session.Close()
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	session.DB(DBNAME).C(DOCNAME).RemoveId(oid)
	// Write status
	return "OK"
}
