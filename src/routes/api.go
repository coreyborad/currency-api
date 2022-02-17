package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterAPI RegisterAPI
func (r *Router) RegisterAPI(api *gin.RouterGroup) error {
	v1 := api.Group("/v1")
	{
		v1.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "API v1",
			})
		})
		v1.GET("currency", r.Currency.CalcCurrency)
	}

	return nil
}
