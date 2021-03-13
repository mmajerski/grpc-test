package utils

import (
	"context"
	"errors"

	"github.com/userq11/grpc-test/models"
	"github.com/userq11/grpc-test/repos"
)

type key string

const (
	globalRepoKey key = "globalRepo"
	userKey       key = "user"
)

func GetGlobalRepoFromContext(ctx context.Context) (repos.GlobalRepository, error) {
	r, ok := ctx.Value(globalRepoKey).(repos.GlobalRepository)
	if ok {
		return r, nil
	}

	return nil, errors.New("missing global repo in context")
}

func SetGlobalRepoOnContext(ctx context.Context, globalRepo repos.GlobalRepository) context.Context {
	return context.WithValue(ctx, globalRepoKey, globalRepo)
}

func SetUserOnContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}
