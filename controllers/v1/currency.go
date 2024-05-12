package v1

import (
	"Asiayo/constant"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get the currency exchange
// @Produce application/json
// @Param route query string true "source"
// @Param route query string true "target"
// @Param route query int true "amount"
// @Success 200 {object} constant.Response
// @Router /api/v1/currency/exchange [get]
func CurrencyExchange(c *gin.Context) {
	var params struct {
		Source string `form:"source" binding:"required"`
		Target string `form:"target" binding:"required"`
		Amount string `form:"amount" binding:"required"`
	}
	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, err.Error())
		return
	}

	fmt.Println(params)

	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, nil)
}
