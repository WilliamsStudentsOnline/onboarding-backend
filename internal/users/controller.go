package users

import (
	"github.com/gin-gonic/gin"
)

// showUser godoc
// @Summary lists the value associated with user from fakedatabase
// @Description Parse the input path param name, query it in fakedatabase, and then write the responses based on whether user exists in database
// @Produce		json
// @Param		name	path	string	true	"name of the user"
// @Success	200	{object}	model.UserInfo
// @Failure	404	{object}	model.UserInfo
// @Router /api/v1/user/{name} [get]
func showUser(c *gin.Context) {
	// TODO: write this function according to the description
}
