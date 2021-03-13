package auth

import (
	"context"

	grpc "github.com/userq11/grpc-test/protobufs"
	"github.com/userq11/grpc-test/utils"
)

type grpcHandler struct{}

func GetRoutes() *grpcHandler {
	return &grpcHandler{}
}

func (h *grpcHandler) Login(ctx context.Context, req *grpc.LoginRequest) (*grpc.LoginReply, error) {
	res := new(grpc.LoginReply)

	// Test on react client
	// res.Token = "asdasdasdasdasdasdsd"
	// return res, nil

	if err := grpc.Validate(req); err != nil {
		return nil, err
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	usersRepo := globalRepo.Users()
	authRepo := globalRepo.Auth()

	user, err := usersRepo.FindByEmail(req.GetEmail())
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(req.GetPassword()); err != nil {
		return nil, err
	}

	claims := authRepo.GetNewClaims(user.Email, map[string]interface{}{
		"user": user,
	})

	tok, err := authRepo.GetSignedToken(claims)
	if err != nil {
		return nil, err
	}

	res.Token = tok

	return res, nil
}
