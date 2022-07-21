package users_mongo

import (
	"context"
	"go-db-segregation/business/model"
	"go-db-segregation/database/mongo/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	usersColl *mongo.Collection
}

func New(db *mongo.Client) *DB {
	return &DB{
		usersColl: db.Database("db").Collection("users"),
	}
}

func (db *DB) GetUsers() (*[]model.User, error) {
	res := []model.User{}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := db.usersColl.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result entities.User
		err = cursor.Decode(&result)
		if err != nil {
			return nil, err
		}

		res = append(res, model.User{
			ID:        result.ID,
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Age:       result.Age,
		})
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return &res, nil
}
