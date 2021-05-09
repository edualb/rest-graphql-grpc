package models

import (
	"context"

	"github.com/edualb/rest-graphql-grpc/server/graphql/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const userCollection = "col_user"

type UsersDB interface {
	CreateUser(u Users) error
	FindUserByID(id string) (*Users, error)
	UpdateUser(id string, u Users) error
	DeleteUserByID(id string) error
}

type UsersDBImpl struct{}

type Users struct {
	Id        primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"name" bson:"name"`
	FlightIDs []primitive.ObjectID `json:"flight_ids" bson:"flight_ids,omitempty"`
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

func (udb UsersDBImpl) FindUserByID(id string) (*Users, error) {
	result := Users{}

	primitiveID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: primitiveID}}

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

func (udb UsersDBImpl) UpdateUser(id string, u Users) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: primitiveID}}

	updater := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "name",
					Value: u.Name,
				},
				primitive.E{
					Key:   "flight_ids",
					Value: u.FlightIDs,
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

func (udb UsersDBImpl) DeleteUserByID(id string) error {
	primitiveID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: primitiveID}}

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
