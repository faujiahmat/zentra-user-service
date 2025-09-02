package cache

import (
	"context"

	"github.com/faujiahmat/zentra-user-service/src/model/entity"
)

type User interface {
	Cache(ctx context.Context, user *entity.User)
	FindByEmail(ctx context.Context, email string) *entity.User
	DeleteByEmail(ctx context.Context, email string)
}
