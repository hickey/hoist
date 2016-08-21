package store

import (
	"github.com/boltdb/bolt"
	"strings"
)

func StartupList() []string {
	return strings.Split(GetValue(prefsBucket, startupKey), arraySeperator)
}

func CreateStartup(name string, container string) error {
	addToStartupList(name)

	// Create the bucket to hold the startup information
	err := SaveValue(name, "container", container)
	if err != nil {
		log.Errorf("Problem creating bucket: %s / %s\n", name, err)
		return err
	}

	// Add the various items that can be stored
	for _, item := range []string{"env", "label", "volume", "port", "link"} {
		err = SaveValue(name, item, "")
		if err != nil {
			log.Errorf("Problem creating item: %s / %s\n", item, err)
			return err
		}
	}

	return nil
}

func DestroyStartup(name string) error {
	removeFromStartupList(name)
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(name))
	})
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
	var startupNames []string

	for _, container := range StartupList() {
		if container != name {
			startupNames = append(startupNames, container)
		}
	}

	err := SaveValue(prefsBucket, startupKey, strings.Join(startupNames, arraySeperator))
	if err != nil {
		log.Error(err)
	}
}
