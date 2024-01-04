package cmd

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/exp/slices"
)

/**** Helpers ****/
const HELPSTRING = "usage: tr string1 string2\n"

func isValidRangePattern(r string) (bool, error) {
	return regexp.MatchString("^.-^|.", r)
}

func createRangeSubstitutions(r1 []rune, r2 []rune) map[rune]rune {
	substitutions := make(map[rune]rune)
	idx := 0
	n1, n2 := len(r1), len(r2)

	for idx = 0; idx < n2 && idx < n1; idx++ {
		substitutions[r1[idx]] = r2[idx]
	}

	/* If len(r1) > len(r2), last character of r2 is duplicated */
	for ; idx < n1; idx++ {
		substitutions[r1[idx]] = r2[n2-1]
	}

	return substitutions
}

/* Assuming a valid pattern matching regexp.MatchString("^.-^|.", r) */
func createRangeFromPattern(p string) []rune {
	runes := []rune{}
	start, _ := utf8.DecodeRuneInString(p)
	end, _ := utf8.DecodeLastRuneInString(p)

	isReverse := end < start
	if isReverse {
		temp := start
		start = end
		end = temp
	}

	runes = append(runes, start)
	for i := start; i <= end; i++ {
		runes = append(runes, i)
	}

	if isReverse {
		slices.Reverse(runes)
	}

	return runes
}

func substitute(s string, m map[rune]rune) string {
	var res strings.Builder

	for i := 0; i < len(s); i++ {
		b := rune(s[i])
		replacement, inMap := m[b]

		if !inMap {
			res.WriteRune(b)
			continue
		}

		res.WriteRune(replacement)
	}

	return res.String()
}
