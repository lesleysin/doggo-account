package base

import (
	"context"

	"github.com/DogGoOrg/doggo-account/internal/handlers"
	"github.com/DogGoOrg/doggo-account/proto/proto_services/Account"
	"gorm.io/gorm"
)

type Server struct {
	Account.UnimplementedAccountServer
	Database *gorm.DB
}

func (r *Server) Ping(ctx context.Context, in *Account.PingRequest) (*Account.PingResponse, error) {
	return handlers.PingHandler(ctx, in)
}

func (r *Server) GetAccountById(ctx context.Context, in *Account.GetAccountRequest) (*Account.GetAccountResponse, error) {
	//TODO - implement me
	return &Account.GetAccountResponse{
		Id:   "122",
		Info: "No info",
	}, nil
}

func (r *Server) Login(ctx context.Context, in *Account.LoginRequest) (*Account.LoginResponse, error) {
	return handlers.LoginHandler(ctx, in, r.Database)
}

func (r *Server) Logout(ctx context.Context, in *Account.LogoutRequest) (*Account.LogoutResponse, error) {
	//TODO - implement me
	return &Account.LogoutResponse{
		Status: "OK",
	}, nil
}

func (r *Server) Refresh(ctx context.Context, in *Account.RefreshRequest) (*Account.RefreshResponse, error) {
	//TODO - implement me
	return &Account.RefreshResponse{
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}

func (r *Server) Register(ctx context.Context, in *Account.RegisterRequest) (*Account.RegisterResponse, error) {
	return handlers.RegisterHandler(ctx, in, r.Database)
}
