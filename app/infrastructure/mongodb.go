package infrastructure

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoDBInfra : Infrastructure layer to incorporate all the MongoDB CRUD
type MongoDBInfra struct{}

//MongoDBOptions : Additional params for Mongodb queries
type MongoDBOptions struct {
	Skip  int64
	Limit int64
}

var _db = DB{}

//FindMany : Generic query to find multiple data from MongoDb
func (mdb *MongoDBInfra) FindMany(collection string, filters interface{}, result interface{}, mongoOptions ...MongoDBOptions) error {
	var err error
	var cursor *mongo.Cursor
	client := _db.GetMongoDB()
	if client != nil {
		findOptions := options.Find()
		if mongoOptions != nil && len(mongoOptions) > 0 {
			findOptions.SetLimit(mongoOptions[0].Limit)
			findOptions.SetSkip(mongoOptions[0].Skip)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		cursor, err = client.Collection(collection).Find(ctx, filters, findOptions)
		if err == nil {
			defer cursor.Close(ctx)
			err = cursor.All(ctx, result)
		}
	}
	return err
}
