package helpers

import (
	"math"
)

func AuthorisationFailure(number int64, lastFour int) bool {
	powerOfEleven := math.Pow10(11)
	lastFourDigits := int(number) % int(powerOfEleven)
	if lastFourDigits == lastFour {
		return true
	} else {
		return false
	}
}
