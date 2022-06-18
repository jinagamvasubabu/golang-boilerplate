package adapters

import (
	"time"

	Logger "github.com/jinagamvasubabu/golang-boilerplate/adapters/logger"
	"github.com/jinagamvasubabu/golang-boilerplate/config"
	mgo "gopkg.in/mgo.v2"
)

var mongoDb *mgo.Database

func InitMongoDatabase() (*mgo.Database, error) {
	// to connect with MongoDB
	cfg := config.GetConfig()
	sess, err := mgo.Dial(cfg.MongoHost)
	if err != nil {
		Logger.Errorf("err:=%s", err.Error())
		time.Sleep(10 * time.Second)
		InitMongoDatabase()
	}

	sess.SetMode(mgo.Monotonic, true)
	mongoDb = sess.DB(cfg.DB)

	return mongoDb, nil
}

func GetConnection() *mgo.Database {
	return mongoDb
}

func SetConnection(conn *mgo.Database) {
	mongoDb = conn
}
