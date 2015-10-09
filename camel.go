package strcase

import (
	"strings"
)

// Converts a string to CamelCase
func ToCamel(s string) string {
	s = strings.Trim(s, " ")
	n := ""
	capNext := true
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_'  || v == ' ' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}
