package service

import (
	"context"

	"github.com/edualb/rest-graphql-grpc/server/grpc/models"
	"github.com/edualb/rest-graphql-grpc/server/grpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error)
	FindUser(ctx context.Context, req *pb.FindUserRequest) (*pb.FindUserResponse, error)
	UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
}

type UserServiceImpl struct {
	db models.UsersDB
}

func NewUserService() UserService {
	return &UserServiceImpl{
		db: models.NewUsersDB(),
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	reqUser := req.GetUser()

	ids := []primitive.ObjectID{}
	for _, id := range reqUser.FlightIds {
		primitiveID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot process uuid: %s", id)
		}
		ids = append(ids, primitiveID)
	}

	user := models.Users{
		Name:      reqUser.Name,
		FlightIDs: ids,
	}

	err := s.db.CreateUser(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot save user: %v", user)
	}

	return &pb.CreateUserResponse{Message: "Success"}, nil
}

func (s *UserServiceImpl) FindUser(ctx context.Context, req *pb.FindUserRequest) (*pb.FindUserResponse, error) {
	reqId := req.GetId()

	user, err := s.db.FindUserByID(reqId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot not found user: %s", reqId)
	}

	ids := []string{}
	for _, id := range user.FlightIDs {
		stringId := id.String()
		ids = append(ids, stringId)
	}

	res := &pb.FindUserResponse{
		Message: "Found",
		User: &pb.User{
			Id:        user.Id.String(),
			Name:      user.Name,
			FlightIds: ids,
		},
	}

	return res, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	reqId := req.GetId()
	reqUser := req.GetUser()

	ids := []primitive.ObjectID{}
	for _, id := range reqUser.FlightIds {
		primitiveID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "cannot process uuid: %s", id)
		}
		ids = append(ids, primitiveID)
	}

	user := models.Users{
		Name:      reqUser.Name,
		FlightIDs: ids,
	}

	err := s.db.UpdateUser(reqId, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot update user: %v", user)
	}

	return &pb.UpdateUserResponse{Message: "Success"}, nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	reqId := req.GetId()

	err := s.db.DeleteUserByID(reqId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot not found user: %s", reqId)
	}

	return &pb.DeleteUserResponse{Message: "User deleted"}, nil
}
