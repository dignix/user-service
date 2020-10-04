package service

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/iam-solutions/user-service/api/v1/pb"
	"github.com/iam-solutions/user-service/internal/app/domain"
)

type userServiceServer struct {}

func NewUserServiceServer() pb.UserServiceServer {
	return &userServiceServer{}
}

func (u *userServiceServer) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	user := &domain.User{
		ID:            1,
		FirstName:     "fsd",
		MiddleName:    "kers",
		LastName:      "arjes",
		Email:         "rweja@asrf.com",
		Password:      "392845qjrn-",
		Avatar:        "/aers/aerkj/r2.jpeg",
		StatusID:      1,
		VerifiedAt:    getTimestamp(time.Now()),
		CreatedAt:     getTimestamp(time.Now()),
		UpdatedAt:     getTimestamp(time.Now()),
	}

	res := &pb.GetResponse{
		User: &pb.User{
			Id:         user.ID,
			FirstName:  user.FirstName,
			MiddleName: user.MiddleName,
			LastName:   user.LastName,
			Email:      user.Email,
			Avatar:     user.Avatar,
			StatusId:   user.StatusID,
			VerifiedAt: user.VerifiedAt,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		},
	}

	return res, nil
}

func (u *userServiceServer) GetAll(request *pb.GetAllRequest, stream pb.UserService_GetAllServer) error {
	users := []*domain.User{
		{
			ID:            1,
			FirstName:     "fsd",
			MiddleName:    "kers",
			LastName:      "arjes",
			Email:         "rweja@asrf.com",
			Password:      "392845qjrn-",
			Avatar:        "/aers/aerkj/r2.jpeg",
			StatusID:      1,
			VerifiedAt:    getTimestamp(time.Now()),
			CreatedAt:     getTimestamp(time.Now()),
			UpdatedAt:     getTimestamp(time.Now()),
		},
		{
			ID:            2,
			FirstName:     "dfk",
			MiddleName:    "aoaejr",
			LastName:      "o32wpjr",
			Email:         "aelsjr@esj.com",
			Password:      "skaerj2",
			Avatar:        "/aers/aerkj/r3.jpeg",
			RememberToken: "",
			StatusID:      2,
			VerifiedAt:    getTimestamp(time.Now()),
			CreatedAt:     getTimestamp(time.Now()),
			UpdatedAt:     getTimestamp(time.Now()),
		},
	}

	for _, user := range users {
		err := stream.Send(&pb.GetAllResponse{
			User: &pb.User{
				Id:            user.ID,
				FirstName:     user.FirstName,
				MiddleName:    user.MiddleName,
				LastName:      user.LastName,
				Email:         user.Email,
				Password:      user.Password,
				Avatar:        user.Avatar,
				StatusId:      user.StatusID,
				VerifiedAt:    user.VerifiedAt,
				CreatedAt:     user.CreatedAt,
				UpdatedAt:     user.UpdatedAt,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func getTimestamp(t time.Time) *timestamp.Timestamp {
	tspb, err := ptypes.TimestampProto(t)
	if err != nil {
		log.Fatalln(err)
	}
	return tspb
}
