package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/itisnotyourenv/todo-go"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		// notify about server error
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// return 200 OK and json with user id
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
