package users

import (
	"fmt"
	"net/http"

	"github.com/WilliamsStudentsOnline/onboarding-backend/internal/fakedatabase"
	"github.com/WilliamsStudentsOnline/onboarding-backend/model"
	"github.com/gin-gonic/gin"
)

// TODO: finish this function according to description
// showUser godoc
// @Summary lists the value associated with user from fakedatabase
// @Description Parse the input path param name, query it in fakedatabase, and then write the responses based on whether user exists in database
// @Produce		json
// @Param		name	path	string	true	"name of the user"
// @Success	200	{object}	model.UserInfo
// @Failure	404	{object}	model.UserInfo
// @Router /api/v1/user/{name} [get]
func showUser(c *gin.Context) {
	user := c.Params.ByName("name")
	value, err := fakedatabase.Query(user)

	// TODO: Delete the line below when done (used to evade unused variable errors)
	fmt.Printf("%v\n", value)

	if err != nil {
		c.JSON(http.StatusNotFound, model.UserInfo{User: user, Color: "Unknown"})
		return
	} else {
		c.JSON(http.StatusOK, model.UserInfo{User: user, Color: "DO LATER"})
		return
	}

	// TODO: If there is no error, repond with a correct JSON that looks like
	// {
	// 	  "user": "Ye",
	//    "color": "Blue"
	// }
}
