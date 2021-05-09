package models

import (
	"context"

	"github.com/edualb/rest-graphql-grpc/server/graphql/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const flightCollection = "col_flights"

type FlightsDB interface {
	CreateFlight(f Flights) error
	UpdateFlightByID(id string, f Flights) error
	FindFlightByID(id string) (*Flights, error)
}

type FlightsDBImpl struct{}

type Flights struct {
	Id           primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	From         string               `json:"from" bson:"from"`
	To           string               `json:"to" bson:"to"`
	PassengerIDs []primitive.ObjectID `json:"passenger_ids,omitempty" bson:"passenger_ids,omitempty"`
}

func NewFlightsDB() FlightsDB {
	return &FlightsDBImpl{}
}

func (udb FlightsDBImpl) CreateFlight(f Flights) error {
	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database((db.DB)).Collection(flightCollection)

	_, err = collection.InsertOne(context.TODO(), f)
	if err != nil {
		return err
	}
	return nil
}

func (udb FlightsDBImpl) FindFlightByID(id string) (*Flights, error) {
	result := Flights{}

	primitiveID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: primitiveID}}

	client, err := db.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database((db.DB)).Collection(flightCollection)

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (udb FlightsDBImpl) UpdateFlightByID(id string, f Flights) error {
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
					Key:   "from",
					Value: f.From,
				},
				primitive.E{
					Key:   "to",
					Value: f.To,
				},
				primitive.E{
					Key:   "passenger_ids",
					Value: f.PassengerIDs,
				},
			},
		},
	}

	client, err := db.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database((db.DB)).Collection(flightCollection)

	_, err = collection.UpdateOne(context.TODO(), filter, updater)

	if err != nil {
		return err
	}

	return nil
}
