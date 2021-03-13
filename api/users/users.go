package users

import (
	"context"

	"github.com/userq11/grpc-test/models"
	grpc "github.com/userq11/grpc-test/protobufs"
	"github.com/userq11/grpc-test/utils"
)

type grpcHandler struct{}

func GetRoutes() grpc.UsersServer {
	return &grpcHandler{}
}

func (h *grpcHandler) Create(ctx context.Context, req *grpc.CreateUserRequest) (*grpc.UserReply, error) {
	res := new(grpc.UserReply)

	if err := grpc.Validate(req); err != nil {
		return nil, err
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	newUser, err := models.NewUser(&models.TempUser{
		Email:           req.GetNewUser().GetEmail(),
		Password:        req.GetNewUser().GetPassword(),
		ConfirmPassword: req.GetNewUser().GetConfirmPassword(),
		FirstName:       req.GetNewUser().GetFirstName(),
		LastName:        req.GetNewUser().GetLastName(),
	})
	if err != nil {
		return nil, err
	}

	if err := globalRepo.Users().Create(newUser); err != nil {
		return nil, err
	}

	res.User = newUser.ToProtobuf()

	return res, nil
}

func (h *grpcHandler) FindByID(ctx context.Context, req *grpc.FindByIDRequest) (*grpc.UserReply, error) {
	res := new(grpc.UserReply)

	if err := grpc.Validate(req); err != nil {
		return nil, err
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := globalRepo.Users().FindByID(req.GetId())
	if err != nil {
		return nil, err
	}

	res.User = user.ToProtobuf()

	return res, nil
}

func (h *grpcHandler) FindByEmail(ctx context.Context, req *grpc.FindByEmailRequest) (*grpc.UserReply, error) {
	res := new(grpc.UserReply)

	if err := grpc.Validate(req); err != nil {
		return nil, err
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	user, err := globalRepo.Users().FindByEmail(req.GetEmail())
	if err != nil {
		return nil, err
	}

	res.User = user.ToProtobuf()

	return res, nil
}

func (h *grpcHandler) Update(ctx context.Context, req *grpc.UpdateUserRequest) (*grpc.UserReply, error) {
	res := new(grpc.UserReply)

	if err := grpc.Validate(req); err != nil {
		return nil, err
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	usersRepo := globalRepo.Users()

	user, err := usersRepo.FindByID(req.GetId())
	if err != nil {
		return nil, err
	}

	if err = user.SetPassword(req.GetNewPassword()); err != nil {
		return nil, err
	}

	if err = usersRepo.Update(user); err != nil {
		return nil, err
	}

	return res, nil
}

func (h *grpcHandler) Delete(ctx context.Context, req *grpc.DeleteUserRequest) (*grpc.UserReply, error) {
	res := new(grpc.UserReply)

	if err := grpc.Validate(req); err != nil {
		return nil, err
	}

	globalRepo, err := utils.GetGlobalRepoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	usersRepo := globalRepo.Users()

	user, err := usersRepo.FindByID(req.GetId())
	if err != nil {
		return nil, err
	}

	if err = usersRepo.Delete(user); err != nil {
		return nil, err
	}

	res.User = user.ToProtobuf()

	return res, nil
}
