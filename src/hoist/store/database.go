package store

import (
	"app"
	"fmt"
	"github.com/boltdb/bolt"
	"strings"
)

var db *bolt.DB

var arraySeperator = ";"
var startupKey = "startup_list"
var prefsBucket = "__prefs__"

var log = app.GetLogger()

func Start() {
	var err error

	db, err = bolt.Open("hoist.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Database code started")
}

func Stop() {
	db.Close()
}

func CreateStartup(name string, container string) error {
	addToStartupList(name)
	return SaveValue(name, "container", container)
}

func DestroyStartup(name string) error {
	removeFromStartupList(name)
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(name))
	})
}

func SaveValue(startup string, name string, value string) error {
	fmt.Printf("DB = %s", db)
	return db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(startup))
		if err != nil {
			log.Errorf("create bucket: %s", err)
			return err
		}

		err = bucket.Put([]byte(name), []byte(value))
		if err != nil {
			log.Errorf("Can not store %s=%s in %s", name, value, startup)
			return err
		}
		return nil
	})
}

func GetValue(startup string, name string) string {
	var val string

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(startup))
		if bucket == nil {
			// if the bucket does not exist return nil
			return nil
		}

		val = string(bucket.Get([]byte(name)))
		// if err != nil {
		// 	log.Errorf("Can not find %s in %s", name, startup)
		// 	return err
		// }
		return nil
	})
	return val
}

func DumpStartup(startup string) {
	fmt.Println("Database representation")
	fmt.Println(db.GoString())

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(startup))
		fmt.Printf("%s contents:", startup)
		fmt.Println("==============================")
		bucket.ForEach(func(k, v []byte) error {
			fmt.Printf("    %s = %s", string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func addToStartupList(name string) {
	var startupNames []string

	startupNames = StartupList()
	if startupNames[0] == "" {
		// special case where new list gets empty value in [0]
		startupNames = []string{name}
	} else {
		startupNames = append(startupNames, name)
	}
	err := SaveValue(prefsBucket, startupKey, strings.Join(startupNames, arraySeperator))
	if err != nil {
		log.Error(err)
	}
}

func removeFromStartupList(name string) {

}

func StartupList() []string {
	return strings.Split(GetValue(prefsBucket, startupKey), arraySeperator)
}
