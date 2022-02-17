//go:build wireinject
// +build wireinject

package routes

import (
	"currency/controllers"
	"net/http"

	"github.com/google/wire"
)

var (
	RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))
	RouteSet  = wire.NewSet(Load, RouterSet)
)

// InitRoute InitRoute
func InitRoute() http.Handler {
	wire.Build(
		controllers.CreateCurrencyController,
		RouteSet,
	)

	return nil
}
