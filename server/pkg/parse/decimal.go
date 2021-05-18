package parse

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// IntWithDefault will try to parse string and return value. If parsing fails, it will return default value
func IntWithDefault(value string, defaultValue int) int {
	result, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int(result)
}

// DecimalWithDefault will try to parse string and return decimal value. If parsing fails, it will return default value
func DecimalWithDefault(value string, defaultValue decimal.Decimal) decimal.Decimal {
	result, err := decimal.NewFromString(value)
	if err != nil {
		return defaultValue
	}
	return result
}

// DecimalToFloat maps decimal to float ignoring success boolean
func DecimalToFloat(d decimal.Decimal) float64 {
	v, _ := d.Float64()
	return v
}
