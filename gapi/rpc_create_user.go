package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	db "github.com/muhammadali7768/simplebank/db/sqlc"
	"github.com/muhammadali7768/simplebank/pb"
	"github.com/muhammadali7768/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}
	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}
	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		// Check if it's a foreign key violation
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23503", "23505":
				return nil, status.Errorf(codes.AlreadyExists, "user already exists: %s", err)
			}

		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return rsp, nil
}
