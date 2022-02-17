package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(r IRouter) http.Handler {
	e := gin.New()
	e.Use(gin.Logger())
	gin.SetMode(gin.DebugMode)
	// middleware
	// middleware.RouteMiddleware(e)

	serviceRoute := e.Group("/")
	// api
	api := serviceRoute.Group("/api")
	{
		r.RegisterAPI(api)
	}

	e.NoMethod(NotFound)

	return e
}

// NotFound represents the 404 page.
func NotFound(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, "PAGE NOT FOUND")
}
