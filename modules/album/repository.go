package album

import (
	"fmt"
	"log"
	"os"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}
var SERVER = os.Getenv("MONGO_HOST") + os.Getenv("MONGO_PORT")
var info = &mgo.DialInfo{
        Addrs:    []string{SERVER},
        Timeout:  60 * time.Second,
        Database: os.Getenv("MONGO_DATBASE"),
        Username: os.Getenv("MONGO_USER"),
        Password: os.Getenv("MONGO_PASSWORD"),
    }
// SERVER the DB server
//var SERVER = os.Getenv("MONGO_HOST") + os.Getenv("MONGO_PORT")

// DBNAME the name of the DB instance
var DBNAME = os.Getenv("MONGO_DATBASE")

// DOCNAME the name of the document
var DOCNAME = "albums"


// GetAlbums returns the list of Albums
func (r Repository) GetAlbums() Albums {
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo info:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Albums{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

//get Album
func (r Repository) GetAlbum(id string) Album {
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo info:", err)
	}
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
func (r Repository) AddAlbum(album Album) bool {
	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(album)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// UpdateAlbum updates an Album in the DB (not used for now)
func (r Repository) UpdateAlbum(album Album) bool {
	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(album.ID, album)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// DeleteAlbum deletes an Album (not used for now)
func (r Repository) DeleteAlbum(id string) string {
	session, err := mgo.DialWithInfo(info)
	defer session.Close()
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}
	// Write status
	return "OK"
}
