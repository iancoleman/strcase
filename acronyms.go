package strcase

import (
	"strings"
	"sync"
)

var uppercaseAcronym = sync.Map{}

//"ID": "id",

// ConfigureAcronym allows you to add additional words which will be considered acronyms
func ConfigureAcronym(key, val string) {
	uppercaseAcronym.Store(key, val)
}

func checkAcronym(key string) (string, bool) {
	if a, ok := uppercaseAcronym.Load(key); ok {
		return a.(string), true
	}
	if a, ok := uppercaseAcronym.Load(strings.ToUpper(key)); ok {
		return a.(string), true
	}
	if a, ok := uppercaseAcronym.Load(strings.ToLower(key)); ok {
		return a.(string), true
	}
	return "", false
}

func isAllUpper(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			return false
		}
	}
	return true
}
