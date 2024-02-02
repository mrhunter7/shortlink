package storeRepository

import (
	"context"

	link "github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1"
	rpc "github.com/shortlink-org/shortlink/boundaries/link/metadata/domain/metadata/v1"
	"github.com/shortlink-org/shortlink/pkg/notify"
)

type Repository interface {
	Get(context.Context, string) (*rpc.Meta, error)
	Add(context.Context, *rpc.Meta) error
}

// Store abstract type
type MetaStore struct {
	Store Repository
	notify.Subscriber[link.Link]
	typeStore string
}
