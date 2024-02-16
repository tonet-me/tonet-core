package cardmongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (d DB) IsCardExistByName(ctx context.Context, name string) (bool, error) {
	filter := bson.D{{"name", name}}
	counted, err := d.collection.CountDocuments(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) { //instead of if err == mongo.ErrNoDocuments
			return false, nil
		}

		return false, err
	}

	if counted != 0 {
		return true, nil
	}
	return false, nil
}