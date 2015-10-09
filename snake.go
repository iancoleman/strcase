// Package strcase converts strings to snake_case or CamelCase
package strcase

import (
	"strings"
)

// Converts a string to snake_case
func ToSnake(s string) string {
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		if i > 0 && v >= 'A' && v <= 'Z' && n[len(n)-1] != '_' {
			n += "_" + string(v)
		} else if v == ' ' {
			n += "_"
		} else {
			n = n + string(v)
		}
	}
	n = strings.ToLower(n)
	return n
}
