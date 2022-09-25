package admin

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r gin.IRouter) {
	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/api/v1/admin/add \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	r.POST("add", addUser)
}
