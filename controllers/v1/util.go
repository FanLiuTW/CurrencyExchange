package v1

import (
	"fmt"
	"strconv"
)

func ValidateFloatDecimal(amount float64) (error, bool) {
	trim := fmt.Sprintf("%.2f", amount)
	value, err := strconv.ParseFloat(trim, 32)
	if err != nil {
		return err, false
	}
	if value == amount {
		return nil, true
	}
	return nil, false
}
