package admin

import (
	"github.com/gin-gonic/gin"
)

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
	// TODO: write this function according to description
}
