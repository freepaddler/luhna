// Package luhna implements tools for "Luhn algorithm"
package luhna

import (
	"fmt"
	"math/rand"
	"strings"
)

// runeIsNotDigit returns false when c is not a decimal digit characters
func runeIsNotDigit(c rune) bool {
	return c < '0' || c > '9'
}

// IsDigitsString returns true id s consists only of decimal digit characters
func IsDigitsString(s string) bool {
	if s == "" {
		return false
	}
	return strings.IndexFunc(s, runeIsNotDigit) == -1
}

// Validate returns true if string represents valid Luhn number
func Validate(s string) bool {
	if !IsDigitsString(s) {
		return false
	}
	sum := 0
	marker := len(s)%2 == 0
	for i := 0; i < len(s); i++ {
		if marker {
			mult := int(s[i]-'0') * 2
			if mult > 9 {
				mult -= 9
			}
			sum += mult
		} else {
			sum += int(s[i] - '0')
		}
		marker = !marker
	}
	return sum%10 == 0
}

// Generate returns new Luhn string with predefined prefix and length
func Generate(prefix string, l int) (s string) {
	if len(prefix) >= l {
		return
	}
	if prefix != "" && !IsDigitsString(prefix) {
		return
	}
	s = prefix
	sum := 0
	marker := l%2 == 0
	for i := 0; i < l-1; i++ {
		if i >= len(prefix) {
			// add random number
			s += fmt.Sprintf("%c", (rand.Int()%9)+'0')
		}
		if marker {
			mult := int(s[i]-'0') * 2
			if mult > 9 {
				mult -= 9
			}
			sum += mult
		} else {
			sum += int(s[i] - '0')
		}
		marker = !marker
	}
	if rest := sum % 10; rest == 0 {
		s += fmt.Sprintf("%c", '0')
	} else {
		s += fmt.Sprintf("%c", 10-rest+'0')
	}
	return
}
