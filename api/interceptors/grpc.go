package interceptors

import (
	"context"
	"errors"
	"reflect"

	"github.com/userq11/grpc-test/repos"
	"github.com/userq11/grpc-test/utils"
	"google.golang.org/grpc"
	"xorm.io/xorm"
)

func GlobalRepoInjector(db *xorm.Engine) grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		globalRepo := repos.GlobalRepo(db)
		newCtx := utils.SetGlobalRepoOnContext(ctx, globalRepo)

		v := reflect.Indirect(reflect.ValueOf(req))
		vField := reflect.Indirect(v.FieldByName("JWT"))

		if !vField.IsValid() {
			return handler(newCtx, req)
		}

		jwtToken := vField.String()
		if len(jwtToken) == 0 {
			return nil, errors.New("unauthorized")
		}

		user, err := globalRepo.Auth().GetDataFromToken(jwtToken)
		if err != nil {
			return nil, errors.New("unauthorized")
		}

		newCtx = utils.SetUserOnContext(newCtx, user)

		// Make request
		res, err := handler(newCtx, req)

		return res, err
	})
}
