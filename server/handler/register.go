package handler

import (
	"context"
	"net/http"

	pb "github.com/elsumanta/grpcserver/grpc"
	"github.com/elsumanta/grpcserver/server/model"
)

// integrate widget
func (s *Server) Register(ctx context.Context, in *pb.RegRequest) (*pb.RegResponse, error) {
	_, err := s.repo.Register(ctx, model.Register{
		FirstName:   in.GetFirstName(),
		LastName:    in.GetLastName(),
		Address:     in.GetAddress(),
		PhoneNumber: in.GetPhoneNumber(),
	})

	if err != nil {
		return &pb.RegResponse{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}, err
	}
	return &pb.RegResponse{
		HttpStatus: http.StatusCreated,
		Message:    "successfully registered",
	}, err
}
