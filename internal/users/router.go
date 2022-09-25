package users

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r gin.IRouter) {
	// Get user value
	r.GET("/user/:name", showUser)
}
