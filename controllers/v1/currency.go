package v1

import (
	"Asiayo/constant"
	"Asiayo/service"
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
		Source string  `form:"source" binding:"required,oneof=TWD JPY USD"`
		Target string  `form:"target" binding:"required,oneof=TWD JPY USD"`
		Amount float64 `form:"amount" binding:"required,gt=0"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.INVALID_PARAMS, err.Error())
		return
	}

	// currency rates
	staticRates := map[string]map[string]float64{
		"TWD": {
			"TWD": 1,
			"JPY": 3.669,
			"USD": 0.03281,
		},
		"JPY": {
			"TWD": 0.26956,
			"JPY": 1,
			"USD": 0.00885,
		},
		"USD": {
			"TWD": 30.444,
			"JPY": 111.801,
			"USD": 1,
		},
	}

	// create CurrencyExchangeService instance
	exchangeService := service.NewCurrencyExchangeService(staticRates)

	// convert the currency from request
	convertedAmount, err := exchangeService.Convert(params.Source, params.Target, params.Amount)
	if err != nil {
		constant.ResponseWithData(c, http.StatusOK, constant.ERROR, err.Error())
		return
	}

	constant.ResponseWithData(c, http.StatusOK, constant.SUCCESS, convertedAmount)
}
