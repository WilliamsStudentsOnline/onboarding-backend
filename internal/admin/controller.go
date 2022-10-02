package admin

import (
	"fmt"
	"net/http"

	"github.com/WilliamsStudentsOnline/onboarding-backend/model"
	"github.com/gin-gonic/gin"
)

// TODO: finish this function according to description
// showUser godoc
// @Summary edits the currently logged in user in fakedatabse
// @Description Parse the input body json, add/edit it in fakedatabase, and then write the responses
// @Security	BasicAuth
// @Accept		json
// @Produce		json
// @Param	equest	body	model.UserInfo	true	"user value to edit"
// @Success	200	{object}	model.EditResponse
// @Failure	500	{object}	model.EditResponse
// @Router /api/v1/admin/edit/ [POST]
func editUser(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	// TODO: Delete the line below when done (used to evade unused variable errors)
	fmt.Printf("%v\n", user)

	// Parse JSON
	var inputJson model.UserInfo

	if c.Bind(&inputJson) == nil {
		// TODO: Make a call to the Insert function of the database and store the return as "err"
		var err error
		if err != nil {
			// TODO: Respond with a StatusOK and a EditResponse JSON
			c.JSON(http.StatusOK, inputJson)
			return
		}
		c.JSON(http.StatusInternalServerError, model.EditResponse{Status: err.Error()})
	}
}
