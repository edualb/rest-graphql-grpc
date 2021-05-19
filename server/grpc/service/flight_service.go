package service

import (
	"context"

	"github.com/edualb/rest-graphql-grpc/server/grpc/models"
	"github.com/edualb/rest-graphql-grpc/server/grpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FlightService interface {
	CreateFlight(ctx context.Context, req *pb.CreateFlightRequest) (*pb.CreateFlightResponse, error)
	FindFlight(ctx context.Context, req *pb.FindFlightRequest) (*pb.FindFlightResponse, error)
	UpdateFlight(ctx context.Context, req *pb.UpdateFlightRequest) (*pb.UpdateFlightResponse, error)
}

type FlightServiceImpl struct {
	db models.FlightsDB
}

func NewFlightService() FlightService {
	return &FlightServiceImpl{
		db: models.NewFlightsDB(),
	}
}

func (s *FlightServiceImpl) CreateFlight(ctx context.Context, req *pb.CreateFlightRequest) (*pb.CreateFlightResponse, error) {
	reqFlight := req.GetFlight()

	ids := []primitive.ObjectID{}
	for _, id := range reqFlight.PassengerIds {
		primitiveID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot process uuid: %s", id)
		}
		ids = append(ids, primitiveID)
	}

	flight := models.Flights{
		From:         reqFlight.From,
		To:           reqFlight.To,
		PassengerIDs: ids,
	}

	err := s.db.CreateFlight(flight)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot save flight: %v", flight)
	}

	return &pb.CreateFlightResponse{Message: "Success"}, nil
}

func (s *FlightServiceImpl) FindFlight(ctx context.Context, req *pb.FindFlightRequest) (*pb.FindFlightResponse, error) {
	reqId := req.GetId()

	flight, err := s.db.FindFlightByID(reqId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot not found user: %s", reqId)
	}

	ids := []string{}
	for _, id := range flight.PassengerIDs {
		stringId := id.String()
		ids = append(ids, stringId)
	}

	res := &pb.FindFlightResponse{
		Message: "Found",
		Flight: &pb.Flight{
			Id:           flight.Id.String(),
			From:         flight.From,
			To:           flight.To,
			PassengerIds: ids,
		},
	}

	return res, nil
}

func (s *FlightServiceImpl) UpdateFlight(ctx context.Context, req *pb.UpdateFlightRequest) (*pb.UpdateFlightResponse, error) {
	reqId := req.GetId()
	reqFlight := req.GetFlight()

	ids := []primitive.ObjectID{}
	for _, id := range reqFlight.PassengerIds {
		primitiveID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot process uuid: %s", id)
		}
		ids = append(ids, primitiveID)
	}

	flight := models.Flights{
		From:         reqFlight.From,
		To:           reqFlight.To,
		PassengerIDs: ids,
	}

	err := s.db.UpdateFlightByID(reqId, flight)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot update user: %v", flight)
	}

	return &pb.UpdateFlightResponse{Message: "Success"}, nil
}
