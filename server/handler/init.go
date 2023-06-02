package handler

import "context"

type Server struct{}

func Init(ctx context.Context) Server {
	return Server{}
}
