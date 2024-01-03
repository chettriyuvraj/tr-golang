package cmd

import (
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

/**** Helpers ****/
const HELPSTRING = "usage: tr string1 string2\n"

func isValidRangePattern(r string) (bool, error) {
	return regexp.MatchString("^.-^|.", r)
}

func createRangeSubstitutions(r1 []byte, r2 []byte) map[byte]byte {
	substitutions := make(map[byte]byte)
	idx := 0
	n1, n2 := len(r1), len(r2)

	for idx = 0; idx < n2 && idx < n1; idx++ {
		substitutions[r1[idx]] = r2[idx] /* Assuming only ASCII characters */
	}

	/* If len(r1) > len(r2), last character of r2 is duplicated */
	for ; idx < n1; idx++ {
		substitutions[r1[idx]] = r2[n2-1]
	}

	return substitutions
}

/* Assuming a valid pattern matching regexp.MatchString("^.-^|.", r) */
func createRangeFromPattern(p string) []byte {
	r := []byte{}
	start, end := int(p[0]), int(p[0]) /* Get ASCII code */
	if len(p) > 2 {
		end = int(p[2])
	}

	isReverse := end < start
	if isReverse {
		temp := start
		start = end
		end = temp
	}

	for b := start; b <= end; b++ {
		r = append(r, byte(b))
	}

	if isReverse {
		slices.Reverse(r)
	}

	return r
}

func substitute(s string, m map[byte]byte) string {
	var res strings.Builder

	for i := 0; i < len(s); i++ {
		b := s[i]
		replacement, inMap := m[b]

		if !inMap {
			res.WriteByte(b)
			continue
		}

		res.WriteByte(replacement)
	}

	return res.String()
}
