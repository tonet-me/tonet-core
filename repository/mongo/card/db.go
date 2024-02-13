package cardmongo

import (
	"context"
	mongodb "github.com/tonet-me/tonet-core/repository/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DBName   string `koanf:"db_name"`
	CollName string `koanf:"coll_name"`
}

type DB struct {
	collection *mongo.Collection
}

func New(cfg Config, client *mongodb.DB) *DB {
	return &DB{
		collection: initialCollection(cfg, client),
	}
}

func initialCollection(cfg Config, client *mongodb.DB) *mongo.Collection {
	err := client.GetClient().Database(cfg.DBName).CreateCollection(context.TODO(), cfg.CollName)
	if err != nil {
		panic(err)
	}

	cardCollection := client.GetClient().Database(cfg.DBName).Collection(cfg.CollName)
	indexModelUserID := mongo.IndexModel{
		Keys: bson.D{{"user_id", 1}},
	}
	indexModelName := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	}
	indexModelNameAndStatus := mongo.IndexModel{
		Keys: bson.D{{"name", 1}, {"status", 1}},
	}
	indexModelUserIDAndID := mongo.IndexModel{
		Keys: bson.D{{"user_id", 1}, {"_id", 1}},
	}
	indexModelUserIDAndName := mongo.IndexModel{
		Keys: bson.D{{"user_id", 1}, {"name", 1}},
	}

	_, iErr := cardCollection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{indexModelUserID, indexModelName, indexModelNameAndStatus,
		indexModelUserIDAndID, indexModelUserIDAndName})
	if iErr != nil {
		panic(iErr)
	}

	return cardCollection
}