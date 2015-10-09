package strcase

import (
	"testing"
)

func TestToCamel(t *testing.T) {
	cases := [][]string{
		[]string{ "test_case", "TestCase" },
		[]string{ "test", "Test" },
		[]string{ "TestCase", "TestCase" },
		[]string{ " test  case ", "TestCase" },
		[]string{ "", "" },
		[]string{ "many_many_words", "ManyManyWords" },
		[]string{ "AnyKind of_string", "AnyKindOfString" },
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToCamel(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
