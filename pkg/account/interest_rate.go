package account

import (
	"github.com/shopspring/decimal"
)

// CalculateInterestApplicable given a balance, returns the interest applicable to the balance
func CalculateInterestApplicable(balance float64) (decimal.Decimal, error) {
	interestApplicable, err := MapBalanceToApplicableInterest(balance)
	if err != nil {
		return decimal.NewFromFloat(NoOpFloatValue), err
	}

	result := balance * interestApplicable

	// Use 4 decimal places for precision
	interestAmount := decimal.NewFromFloat(result).Round(4)

	// Return in cents, typicall 2 decimal places
	return interestAmount.Round(2), nil
}
