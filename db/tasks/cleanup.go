package tasks

import "github.com/FoxComm/libs/Godeps/_workspace/src/gopkg.in/mgo.v2"

// CleanMongoCollections is a utility that will cleanup collections in Mongo.
// If the collection does not exist, this acts as a no-op.
func CleanMongoCollections(collections ...string) error {
	session, err := mgo.Dial("localhost")
	defer session.Close()

	if err != nil {
		return err
	}

	db := session.DB("social_analytics_test")

	for _, collection := range collections {
		c := db.C(collection)
		if count, err := c.Count(); err != nil {
			return err
		} else if count > 0 {
			if err := c.DropCollection(); err != nil {
				return err
			}
		}
	}

	return nil
}
