package accounting

import (
	"math/big"

	"github.com/cockroachdb/apd"
	"github.com/shopspring/decimal"
)

func formatNumberString(x string, precision int, thousand string, decimalStr string) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatNumber is a base function of the library which formats a number with custom precision and separators.
// FormatNumber supports various types of value by runtime reflection.
// If you don't need runtime type evaluation, please refer to FormatNumberInt or FormatNumberFloat64.
// (supported value types : int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, *big.Rat)
// (also supported value types : decimal.Decimal, *decimal.Decimal *apd.Decimal)
func FormatNumber(value interface{}, precision int, thousand string, decimalStr string) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatNumberInt only supports int value. It is faster than FormatNumber,
// because it does not do any runtime type evaluation.
func FormatNumberInt(x int, precision int, thousand string, decimalStr string) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatNumberFloat64 only supports float64 value.
// It is faster than FormatNumber, because it does not do any runtime type evaluation.
func FormatNumberFloat64(x float64, precision int, thousand string, decimalStr string) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatNumberBigRat only supports *big.Rat value.
// It is faster than FormatNumber, because it does not do any runtime type evaluation.
func FormatNumberBigRat(x *big.Rat, precision int, thousand string, decimalStr string) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatNumberBigDecimal only supports *apd.Decimal value.
// It is faster than FormatNumber, because it does not do any runtime type evaluation.
func FormatNumberBigDecimal(x *apd.Decimal, precision int, thousand string, decimalStr string) string {
	_ = "STUB: not implemented"
	return ""
}

// FormatNumberDecimal only supports decimal.Decimal value.
// It is faster than FormatNumber, because it does not do any runtime type evaluation.
func FormatNumberDecimal(x decimal.Decimal, precision int, thousand string, decimalStr string) string {
	_ = "STUB: not implemented"
	return ""
}
