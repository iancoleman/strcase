// Package strcase converts strings to snake_case or CamelCase
package strcase

import (
	"strings"
)

// Converts a string to snake_case
func ToSnake(s string) string {
	s = addWordBoundariesToNumbers(s)
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		nextIsCapital := false
		if i + 1 < len(s) {
			w := s[i+1]
			nextIsCapital = w >= 'A' && w <= 'Z'
		}
		if i > 0 && v >= 'A' && v <= 'Z' && n[len(n)-1] != '_' && !nextIsCapital {
			// add underscore if next letter is a capital
			n += "_" + string(v)
		} else if v == ' ' {
			// replace spaces with underscores
			n += "_"
		} else {
			n = n + string(v)
		}
	}
	n = strings.ToLower(n)
	return n
}
