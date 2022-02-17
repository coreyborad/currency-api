// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package controllers

import (
	"currency/services"
	"github.com/google/wire"
)

// Injectors from wire.go:

// CreateCurrencyController CreateCurrencyController
func CreateCurrencyController() *CurrencyController {
	currencyService := services.CreateCurrencyService()
	currencyController := NewCurrencyController(currencyService)
	return currencyController
}

// wire.go:

var (
	CurrencyControllerSet = wire.NewSet(NewCurrencyController, services.CreateCurrencyService)
)