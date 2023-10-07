package mongodb

import (
	"context"
	"log"
	"time"

	"github.com/audryus/miniature-octo-tribble/config"
	"github.com/audryus/miniature-octo-tribble/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func New(ctx context.Context) context.Context {
	config := config.Get(ctx)

	ctxMongo, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxMongo, options.Client().ApplyURI(config.Mongo.URL))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.Mongo.DATABASE)

	err = client.Ping(ctxMongo, nil)

	if err != nil {
		log.Fatal(err)
	}

	return context.WithValue(ctx, types.Mongodb, &MongoInstance{
		Client: client,
		Db:     db,
	})

}

func Get(ctx context.Context) *MongoInstance {
	return ctx.Value(types.Mongodb).(*MongoInstance)
}
