package handler

import (
	"form-to-mail-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) sendEmail(c *gin.Context) {
	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.SendEmail(&customer)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "email has been sent successfully")
}
