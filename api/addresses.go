package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListAddresses Get a list of all addresses
// @Summary Get all addresses
// @Tags addresses
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Address
// @Router /addresses [get]
func ListAddresses(c *gin.Context) {
	c.JSON(http.StatusOK, nil) // 200
}
