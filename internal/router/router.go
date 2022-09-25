package router

import (
	"net/http"

	"github.com/WilliamsStudentsOnline/onboarding-backend/internal/admin"
	"github.com/WilliamsStudentsOnline/onboarding-backend/internal/users"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", ping)

	apiV1Group := r.Group("/api/v1")

	// user router: /api/v1/user/*
	userGroup := apiV1Group.Group("/user")
	users.SetupRouter(userGroup)

	// admin router: /api/v1/admin/*
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/admin")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := apiV1Group.Group("/admin",
		gin.BasicAuth(gin.Accounts{
			"foo":  "bar", // user:foo password:bar
			"manu": "123", // user:manu password:123
		}),
	)
	admin.SetupRouter(authorized)

	return r
}

// ping godoc
// @Summary ping example
// @Schemes
// @Description sanity check, return "Hello World"
// @Success 200 {string} Hello World
// @Router /ping [get]
func ping(c *gin.Context) {
	// TODO: Part 1: change the returned text to "Hello World!"
	c.String(http.StatusOK, "pong\n")
}
