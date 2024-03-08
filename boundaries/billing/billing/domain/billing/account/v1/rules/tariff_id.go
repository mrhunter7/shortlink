package rules

import (
	"errors"

	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/domain/billing/account/v1"
)

var ErrTariffIdRequired = errors.New("TariffId is required")

type TariffId struct{}

func NewTariffId() *TariffId {
	return &TariffId{}
}

func (t *TariffId) IsSatisfiedBy(account *v1.Account) error {
	if account.GetTariffId() != "" {
		return nil
	}

	return ErrTariffIdRequired
}
