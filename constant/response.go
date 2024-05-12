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

var msgFlags = map[int]string{
	SUCCESS:        "success",
	INVALID_PARAMS: "invalid params error",
	ERROR:          "fail",
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func ResponseWithData(c *gin.Context, httpCode, respCode int, data interface{}) {
	msg, ok := msgFlags[respCode]
	if !ok {
		msg = msgFlags[ERROR]
	}

	response := Response{
		Code:    respCode,
		Message: msg,
		Data:    data,
	}
	c.JSON(httpCode, response)
}
