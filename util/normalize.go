package util

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

func Normalize(input string) string {
	result := strings.ToLower(input)
	t := transform.Chain(norm.NFC, runes.Remove(runes.In(unicode.Mn)), runes.Remove(runes.Predicate(func(r rune) bool {
		// Remove non-alphanumeric characters, excluding basic Latin letters and numbers
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})))
	normalized, _, err := transform.String(t, result)
	if err != nil {
		return input
	}

	normalized = strings.TrimSpace(normalized)
	return normalized
}
