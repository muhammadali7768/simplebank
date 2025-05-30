package gapi

import (
	"fmt"

	db "github.com/muhammadali7768/simplebank/db/sqlc"
	"github.com/muhammadali7768/simplebank/pb"
	"github.com/muhammadali7768/simplebank/token"
	"github.com/muhammadali7768/simplebank/util"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker : %w", err)
	}
	server := &Server{store: store, config: config, tokenMaker: tokenMaker}

	return server, nil
}
