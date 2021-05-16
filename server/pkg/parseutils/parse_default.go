package parseutils

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// ParseIntWithDefault will try to parse string and return value. If parsing fails, it will return default value
func ParseIntWithDefault(value string, defaultValue int) int {
	result, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int(result)
}

// ParseDecimalWithDefault will try to parse string and return decimal value. If parsing fails, it will return default value
func ParseDecimalWithDefault(value string, defaultValue decimal.Decimal) decimal.Decimal {
	result, err := decimal.NewFromString(value)
	if err != nil {
		return defaultValue
	}
	return result
}
