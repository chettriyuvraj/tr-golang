package cmd

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/exp/slices"
)

/**** Helpers ****/
const HELPSTRING = "usage: tr string1 string2\n"
const (
	RANGE  = "RANGE"
	CLASS  = "CLASS"
	DIRECT = "DIRECT"
)

func findPatternType(p string) (string, error) {
	match, err := regexp.MatchString("^.-.$", p)
	if err != nil {
		return "", err
	}
	if match {
		return RANGE, nil
	}

	match, err = regexp.MatchString("^\\[:.*:\\]$", p)
	if err != nil {
		return "", err
	}
	if match {
		return CLASS, nil
	}

	return DIRECT, nil
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

func createRangeFromPattern(p string) ([]rune, error) {
	var runes []rune

	t, err := findPatternType(p)
	if err != nil {
		return nil, err
	}

	switch t {
	case RANGE:
		runes = rangeAsRunes(p)
	case CLASS:
		runes = classAsRunes(p)
	case DIRECT:
		runes = strAsRunes(p)
	}

	return runes, nil
}

/* Convert a pattern of type 'RANGE' into []rune containing each rune of the range */
func rangeAsRunes(p string) []rune {
	runes := []rune{}
	start, _ := utf8.DecodeRuneInString(p)
	end, _ := utf8.DecodeLastRuneInString(p)

	isReverse := end < start
	if isReverse {
		temp := start
		start = end
		end = temp
	}

	for i := start; i <= end; i++ {
		runes = append(runes, i)
	}

	if isReverse {
		slices.Reverse(runes)
	}

	return runes
}

func strAsRunes(s string) []rune {
	runes := []rune{}
	idx := 0
	for r, size := utf8.DecodeRuneInString(s); idx < len(s); r, size = utf8.DecodeRuneInString(s[idx:]) {
		runes = append(runes, r)
		idx += size
	}
	return runes
}

func classAsRunes(class string) []rune {
	var res []rune

	switch class {
	case "[:alpha:]":
		res = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	case "[:lower:]":
		res = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	case "[:upper:]":
		res = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	}

	return res
}

func substitute(s string, m map[rune]rune) string {
	var res strings.Builder
	idx := 0

	for r, size := utf8.DecodeRuneInString(s); idx < len(s); r, size = utf8.DecodeRuneInString(s[idx:]) {
		replacement, inMap := m[r]

		if !inMap {
			res.WriteRune(r)
			idx += size
			continue
		}

		res.WriteRune(replacement)
		idx += size
	}

	return res.String()
}
