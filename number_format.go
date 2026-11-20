package otto

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

// The helpers below follow ECMA-262 Number.prototype.toFixed,
// toExponential and toPrecision. strconv.FormatFloat can't be used as is:
// it rounds exact ties to even (the spec picks the larger n), pads the
// exponent to two digits and keeps the sign of -0.

var bigTen = big.NewInt(10)

// roundHalfUp returns the integer closest to r * 10^exp, picking the
// larger one on a tie. r must not be negative.
func roundHalfUp(r *big.Rat, exp int) *big.Int {
	scaled := new(big.Rat).Mul(r, powerOfTen(exp))
	// floor((2*num + den) / (2*den)) == floor(scaled + 1/2)
	num := new(big.Int).Lsh(scaled.Num(), 1)
	num.Add(num, scaled.Denom())
	den := new(big.Int).Lsh(scaled.Denom(), 1)
	return num.Quo(num, den)
}

// significantDigits returns the precision digits n and the exponent e such
// that n * 10^(e-precision+1) is as close as possible to value (ties pick
// the larger n). value must be finite and > 0.
func significantDigits(value float64, precision int) (string, int) {
	exact := new(big.Rat).SetFloat64(value)
	// Find e with 10^e <= value < 10^(e+1). Log10 can be off by one near
	// powers of ten, so check against the exact value.
	exponent := int(math.Floor(math.Log10(value)))
	for exact.Cmp(powerOfTen(exponent)) < 0 {
		exponent--
	}
	for exact.Cmp(powerOfTen(exponent+1)) >= 0 {
		exponent++
	}
	digits := roundHalfUp(exact, precision-1-exponent).String()
	if len(digits) > precision {
		// Rounded up to 10^precision, e.g. 9.99 with precision 2.
		return "1" + strings.Repeat("0", precision-1), exponent + 1
	}
	return digits, exponent
}

// powerOfTen returns 10^exp as an exact rational.
func powerOfTen(exp int) *big.Rat {
	if exp >= 0 {
		return new(big.Rat).SetInt(new(big.Int).Exp(bigTen, big.NewInt(int64(exp)), nil))
	}
	return new(big.Rat).SetFrac(big.NewInt(1), new(big.Int).Exp(bigTen, big.NewInt(int64(-exp)), nil))
}

// exponentSuffix formats an exponent the way the spec does: "e+" or "e-"
// followed by the digits, with no leading zeros (e.g. "e+0", "e-7").
func exponentSuffix(exponent int) string {
	if exponent < 0 {
		return "e-" + strconv.Itoa(-exponent)
	}
	return "e+" + strconv.Itoa(exponent)
}

// numberToFixed implements toFixed for a finite value with |value| < 1e21.
func numberToFixed(value float64, fractionDigits int) string {
	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}
	digits := roundHalfUp(new(big.Rat).SetFloat64(value), fractionDigits).String()
	if fractionDigits == 0 {
		return sign + digits
	}
	if len(digits) <= fractionDigits {
		digits = strings.Repeat("0", fractionDigits+1-len(digits)) + digits
	}
	point := len(digits) - fractionDigits
	return sign + digits[:point] + "." + digits[point:]
}

// numberToExponential implements toExponential for a finite value.
// fractionDigits < 0 means undefined: use as many digits as needed to
// represent the value uniquely.
func numberToExponential(value float64, fractionDigits int) string {
	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}
	if fractionDigits < 0 && value != 0 {
		// Shortest round-trip digits. FormatFloat already gives the mantissa
		// in the right shape ("7.71234e+01"); only the exponent needs its
		// leading zero dropped.
		mantissa, exp, _ := strings.Cut(strconv.FormatFloat(value, 'e', -1, 64), "e")
		exponent, _ := strconv.Atoi(exp)
		return sign + mantissa + exponentSuffix(exponent)
	}
	var digits string
	var exponent int
	if value == 0 {
		digits = strings.Repeat("0", max(fractionDigits, 0)+1)
	} else {
		digits, exponent = significantDigits(value, fractionDigits+1)
	}
	if len(digits) > 1 {
		digits = digits[:1] + "." + digits[1:]
	}
	return sign + digits + exponentSuffix(exponent)
}

// numberToPrecision implements toPrecision for a finite value.
func numberToPrecision(value float64, precision int) string {
	sign := ""
	if value < 0 {
		sign = "-"
		value = -value
	}
	var digits string
	var exponent int
	if value == 0 {
		digits = strings.Repeat("0", precision)
	} else {
		digits, exponent = significantDigits(value, precision)
	}
	switch {
	case exponent < -6 || exponent >= precision:
		if precision > 1 {
			digits = digits[:1] + "." + digits[1:]
		}
		return sign + digits + exponentSuffix(exponent)
	case exponent == precision-1:
		return sign + digits
	case exponent >= 0:
		return sign + digits[:exponent+1] + "." + digits[exponent+1:]
	default:
		return sign + "0." + strings.Repeat("0", -(exponent+1)) + digits
	}
}
