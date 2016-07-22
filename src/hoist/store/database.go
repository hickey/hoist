package store

import (
	"app"
	"github.com/boltdb/bolt"
)

var db *bolt.DB
var log = app.GetLogger()

func init() {
	var err error
	//log := app.GetLogger()

	db, err = bolt.Open("hoist.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func Start() {
	log.Debug("Database code started")
}
