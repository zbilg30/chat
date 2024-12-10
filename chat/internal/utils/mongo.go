package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func DecodeAll[T any](ctx context.Context, cursor *mongo.Cursor) ([]T, error) {
	var results []T
	for cursor.Next(ctx) {
		var item T
		if err := cursor.Decode(&item); err != nil {
			continue
		}
		results = append(results, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}