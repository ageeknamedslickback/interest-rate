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
	interestAmount := decimal.NewFromFloat(result)

	return interestAmount.Round(2), nil
}
