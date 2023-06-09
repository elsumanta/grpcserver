package handler

import (
	"context"

	"github.com/elsumanta/grpcserver/server/model"
)

//go:generate moq -out handler_mock.go . Repo
type Repo interface {
	Register(ctx context.Context, req model.Register) (num int, err error)
}
type Server struct {
	repo Repo
}

func Init(ctx context.Context, repo Repo) *Server {
	return &Server{
		repo: repo,
	}
}
