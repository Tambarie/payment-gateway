package helpers

import (
	"log"
	"math"
)

func AuthorisationFailure(number int64, lastFour int) bool {
	powerOfEleven := math.Pow10(11)
	lastFourDigits := int(number) % int(powerOfEleven)
	log.Println(lastFourDigits)
	if lastFourDigits == lastFour {
		return true
	} else {
		return false
	}
}
