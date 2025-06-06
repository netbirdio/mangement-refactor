package routers

import (
	"context"

	"github.com/netbirdio/management-refactor/internals/shared/db"
)

type Manager interface {
	GetNetworkRoutersByNetID(ctx context.Context, tx db.Transaction, lockingStrength db.LockingStrength, accountID, userID, networkID string) ([]*NetworkRouter, error)
	DeleteRouter(ctx context.Context, tx db.Transaction, accountID, userID, networkID, routerID string) error
	DeleteRoutersInNetwork(ctx context.Context, tx db.Transaction, accountID, userID, networkID string) error
}
