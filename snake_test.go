package strcase

import (
	"testing"
)

func TestToSnake(t *testing.T) {
	cases := [][]string{
		[]string{ "testCase", "test_case" },
		[]string{ "TestCase", "test_case" },
		[]string{ "Test Case", "test_case" },
		[]string{ " Test Case", "test_case" },
		[]string{ "Test Case ", "test_case" },
		[]string{ " Test Case ", "test_case" },
		[]string{ "test", "test" },
		[]string{ "test_case", "test_case" },
		[]string{ "Test", "test" },
		[]string{ "", "" },
		[]string{ "ManyManyWords", "many_many_words" },
		[]string{ "manyManyWords", "many_many_words" },
		[]string{ "AnyKind of_string", "any_kind_of_string" },
	}
	for _, i := range cases {
		in := i[0]
		out := i[1]
		result := ToSnake(in)
		if result != out {
			t.Error("'" + result + "' != '" + out + "'")
		}
	}
}
