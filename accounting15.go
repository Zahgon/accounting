//go:build go1.5
// +build go1.5

package accounting

import (
	"math/big"
)

// FormatMoneyBigFloat only supports *big.Float value. It is faster than FormatMoney,
// because it does not do any runtime type evaluation.
func (accounting *Accounting) FormatMoneyBigFloat(value *big.Float) string {
	_ = "STUB: not implemented"
	return ""
}
