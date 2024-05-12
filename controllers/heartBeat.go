package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HeartBeat function
// @Summary HeartBeat
// @Produce plain
// @Success 200 {string} string "Alive"
// @Router /heartBeat [get]
func HeartBeat(c *gin.Context) {
	c.String(http.StatusOK, "Alive")
}
