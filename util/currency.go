package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupporedCurrency returns true if the currency is supported
func IsSupporedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
