//go:build wireinject
// +build wireinject

package services

import (
	"github.com/google/wire"
)

var (
	CurrencyServiceSet = wire.NewSet(NewCurrencyService)
)

// CreateCurrencyService CreateCurrencyService
func CreateCurrencyService() *CurrencyService {
	wire.Build(CurrencyServiceSet)

	return nil
}
