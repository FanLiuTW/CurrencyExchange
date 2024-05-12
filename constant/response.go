package constant

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS        = http.StatusOK
	INVALID_PARAMS = http.StatusBadRequest
	ERROR          = http.StatusInternalServerError
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ResponseWithData(c *gin.Context, httpCode, respCode int, data interface{}) {
	response := Response{
		Code: respCode,
		Data: data,
	}
	c.JSON(httpCode, response)
}
