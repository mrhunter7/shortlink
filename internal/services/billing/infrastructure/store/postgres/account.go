package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	v1 "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/account/v1"
)

type Account struct {
	client *pgxpool.Pool
}

func (a *Account) Init(_ context.Context, store db.DB) error {
	var ok bool

	a.client, ok = store.GetConn().(*pgxpool.Pool)
	if !ok {
		return db.ErrGetConnection
	}

	return nil
}

func (a *Account) Get(ctx context.Context, id string) (*v1.Account, error) {
	panic("implement me")
}

func (a *Account) List(ctx context.Context, filter any) ([]*v1.Account, error) {
	panic("implement me")
}

func (a *Account) Add(ctx context.Context, in *v1.Account) (*v1.Account, error) {
	id := uuid.MustParse(in.GetId())
	userId := uuid.MustParse(in.GetUserId())
	tariffId := uuid.MustParse(in.GetTariffId())

	// query builder
	account := psql.Insert("billing.account").
		Columns("id", "user_id", "tariff_id").
		Values(id, userId, tariffId)

	q, args, err := account.ToSql()
	if err != nil {
		return nil, err
	}

	row := a.client.QueryRow(ctx, q, args...)
	errScan := row.Scan()
	if errors.Is(errScan, pgx.ErrNoRows) {
		return in, nil
	}
	if errScan.Error() != "" {
		return nil, errScan
	}

	return in, nil
}

func (a *Account) Update(ctx context.Context, in *v1.Account) (*v1.Account, error) {
	panic("implement me")
}

func (a *Account) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
