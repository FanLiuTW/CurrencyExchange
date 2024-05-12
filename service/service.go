package service

import (
	"fmt"
	"math"
)

type CurrencyExchangeService struct {
	rates map[string]map[string]float64
}

func NewCurrencyExchangeService(rates map[string]map[string]float64) *CurrencyExchangeService {
	return &CurrencyExchangeService{
		rates: rates,
	}
}

func (s *CurrencyExchangeService) Convert(source, target string, amount float64) (string, error) {
	if _, ok := s.rates[source]; !ok {
		return "", fmt.Errorf("source currency %s not supported", source)
	}
	if _, ok := s.rates[target]; !ok {
		return "", fmt.Errorf("target currency %s not supported", target)
	}
	if amount <= 0 {
		return "", fmt.Errorf("amount must be a positive number")
	}

	if source == target {
		return formatAmount(amount), nil
	}

	rate := s.rates[source][target]
	convertedAmount := amount * rate

	return formatAmount(convertedAmount), nil
}

func formatAmount(amount float64) string {
	roundedAmount := math.Round(amount*100) / 100
	formattedAmount := fmt.Sprintf("%.2f", roundedAmount)
	formattedAmountWithComma := ""
	negative := false

	if formattedAmount[0] == '-' {
		negative = true
		formattedAmount = formattedAmount[1:]
	}

	dotIndex := len(formattedAmount)
	for i, char := range formattedAmount {
		if char == '.' {
			dotIndex = i
			break
		}
	}

	for i, j := dotIndex-1, 0; i >= 0; i, j = i-1, j+1 {
		if j > 0 && j%3 == 0 {
			formattedAmountWithComma = "," + formattedAmountWithComma
		}
		formattedAmountWithComma = string(formattedAmount[i]) + formattedAmountWithComma
	}
	formattedAmountWithComma = formattedAmountWithComma + formattedAmount[dotIndex:]

	if negative {
		formattedAmountWithComma = "-" + formattedAmountWithComma
	}
	return formattedAmountWithComma
}
