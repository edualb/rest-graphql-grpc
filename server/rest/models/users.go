package models

import (
	"context"

	"github.com/edualb/rest-graphql-grpc/server/rest/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const userCollection = "col_user"

type UsersDB interface {
	CreateUser(u Users) error
	FindUserByName(name string) (*Users, error)
	UpdateUser(u Users) error
	DeleteUser(u Users) error
}

type UsersDBImpl struct{}

type Users struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

func NewUsersDB() UsersDB {
	return &UsersDBImpl{}
}

func (udb UsersDBImpl) CreateUser(u Users) error {
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database((db.DB)).Collection(userCollection)

	_, err = collection.InsertOne(context.TODO(), u)
	if err != nil {
		return err
	}
	return nil
}

func (udb UsersDBImpl) FindUserByName(name string) (*Users, error) {
	result := Users{}

	filter := bson.D{primitive.E{Key: "name", Value: name}}

	client, err := db.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database((db.DB)).Collection(userCollection)

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (udb UsersDBImpl) UpdateUser(u Users) error {
	filter := bson.D{primitive.E{Key: "_id", Value: u.Id}}

	updater := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "name",
					Value: u.Name,
				},
			},
		},
	}

	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database((db.DB)).Collection(userCollection)

	_, err = collection.UpdateOne(context.TODO(), filter, updater)

	if err != nil {
		return err
	}

	return nil
}

func (udb UsersDBImpl) DeleteUser(u Users) error {
	filter := bson.D{primitive.E{Key: "_id", Value: u.Id}}

	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database((db.DB)).Collection(userCollection)

	_, err = collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}
