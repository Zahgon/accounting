package accounting

// UnformatNumber takes a string of the number to strip currency info on
// and precision for decimals.
// It pulls the currency descripter from the LocaleInfo map and uses it to return an unformatted value
// based on thous sep and decimal sep
func UnformatNumber(n string, precision int, currency string) string {
	_ = "STUB: not implemented"
	return ""
}

// Remove anything thats not a space, comma, or decimal

// Strip out thousands seperator, whatever it is

// Replace decimal seperator with a decimal at specified precision

func setPrecision(num string, precision int) string { _ = "STUB: not implemented"; return "" }
