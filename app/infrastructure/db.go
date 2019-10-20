package infrastructure

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	mgo "gopkg.in/mgo.v2"
)

//DB  Infrastructure layer to incorporate all the database function
type DB struct{}

var _dbConfig = ConfigService{}

//MongoDB ...
type MongoDB struct {
	MongoDatabase   *mgo.Database
	MongoDBDatabase *mongo.Database
}

var _mongoInstance *MongoDB
var _session *mgo.Session

var _initMongoDBCtx sync.Once
var _mongoClient *mongo.Client

//GetMongoDB ...
func (ds *DB) GetMongoDB() *mongo.Database {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if _mongoClient != nil {
		err = _mongoClient.Ping(ctx, readpref.Primary())
		if err == nil {
			return _mongoInstance.MongoDBDatabase
		}
	}
	_initMongoDBCtx.Do(func() {
		log.Println("Connecting to MongoDB")
		_config := _dbConfig.GetConfiguration()
		_mongoClient, err = mongo.Connect(ctx, &options.ClientOptions{
			Hosts: []string{_config.MongoDBHost},
			Auth: &options.Credential{
				Username:      _config.MongoDBUsername,
				Password:      _config.MongoDBPassword,
				AuthMechanism: "SCRAM-SHA-1",
				AuthSource:    _config.MongoDBDatabase,
			},
		})
		if err == nil {
			err = _mongoClient.Ping(ctx, readpref.Primary())
			if err == nil {
				if _session == nil {
					_mongoInstance = &MongoDB{MongoDBDatabase: _mongoClient.Database(_config.MongoDBDatabase)}
				} else {
					_mongoInstance = &MongoDB{
						MongoDBDatabase: _mongoClient.Database(_config.MongoDBDatabase),
						MongoDatabase:   _session.DB(_config.MongoDBDatabase)}
				}
				log.Println("Connected to Mongo...")
			}
		} else {
			log.Println("Failed to connect to Mongo..." + " " + err.Error())
		}
	})
	return _mongoInstance.MongoDBDatabase
}
