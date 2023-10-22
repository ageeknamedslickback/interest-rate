package account

import "fmt"

const (
	// FirstBand represents a range up to but not inclusive of the figure (1000)
	FirstBand float64 = 1000

	// FirstBandInterest is the interest applicable to the first band balance
	FirstBandInterest float64 = 0.01

	// SecondBand represents a range from & inclusive of the FirstBand up to but not inclusive of the SecondBand figure (5000)
	SecondBand float64 = 5000

	// SecondBandInterest is the interest applicable to the Second band balance
	SecondBandInterest float64 = 0.015

	// ThirdBand represents a range from & inclusive of the SecondBand up to but not inclusive of the ThirdBand figure (10000)
	ThirdBand float64 = 10000

	// ThirdBandInterest is the interest applicable to the Third band balance
	ThirdBandInterest float64 = 0.02

	// LastBand represents a range from & inclusive of the ThirdBand up to but not inclusive of the LastBand figure (50000)
	// This band implicitly handles balances that are also above the specified figure
	LastBand float64 = 50000

	// LastBandInterest is the interest applicable to the Last band balance
	LastBandInterest float64 = 0.025

	// GTELastBandInterest is the interest applicable to the Greater Than or Equal to Last band balance
	GTELastBandInterest float64 = 0.03

	// NoOpInterest represents a no operation interest value
	NoOpInterest float64 = 0.0
)

// InterestRatingTable maps the balance bands to their respective applicable interests
var InterestRatingTable = map[float64]float64{
	FirstBand:  FirstBandInterest,
	SecondBand: SecondBandInterest,
	ThirdBand:  ThirdBandInterest,
	LastBand:   LastBandInterest,
}

// MapBalanceToApplicableInterest returns the interest applicable to a balance's band
func MapBalanceToApplicableInterest(balance float64) (float64, error) {
	if balance > 0 && balance < FirstBand {
		return InterestRatingTable[FirstBand], nil
	}

	if balance >= FirstBand && balance < SecondBand {
		return InterestRatingTable[SecondBand], nil
	}

	if balance >= SecondBand && balance < ThirdBand {
		return InterestRatingTable[ThirdBand], nil
	}

	if balance >= ThirdBand && balance < LastBand {
		return InterestRatingTable[LastBand], nil
	}
	if balance >= LastBand {
		return GTELastBandInterest, nil
	}

	return NoOpInterest, fmt.Errorf("unable to get the interest applicable to a balance of %v", balance)
}
