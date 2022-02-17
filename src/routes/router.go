package routes

import (
	"currency/controllers"

	"github.com/gin-gonic/gin"
)

var _ IRouter = (*Router)(nil)

// IRouter IRouter
type IRouter interface {
	RegisterAPI(api *gin.RouterGroup) error
}

// Router 路由註冊
type Router struct {
	Currency *controllers.CurrencyController
}
