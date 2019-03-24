package db

import (
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)



var SERVER = os.Getenv("MONGO_HOST") + os.Getenv("MONGO_PORT")
var info = &mgo.DialInfo{
	Addrs:    []string{SERVER},
	Timeout:  60 * time.Second,
	Database: os.Getenv("MONGO_DATABASE"),
	Username: os.Getenv("MONGO_USER"),
	Password: os.Getenv("MONGO_PASSWORD"),
}


func Session() *mgo.Session {
	mongoSession, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	return mongoSession;
}
