//go:build wireinject
// +build wireinject

package controllers

import (
	"currency/services"

	"github.com/google/wire"
)

var (
	CurrencyControllerSet = wire.NewSet(NewCurrencyController, services.CreateCurrencyService)
)

// CreateCurrencyController CreateCurrencyController
func CreateCurrencyController() *CurrencyController {
	wire.Build(CurrencyControllerSet)

	return nil
}
