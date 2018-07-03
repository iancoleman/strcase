# strcase
[![Build Status](https://travis-ci.org/Ma124/strcase.svg)](https://travis-ci.org/Ma124/strcase)
[![Coverage](http://gocover.io/_badge/github.com/Ma124/strcase?0)](http://gocover.io/github.com/Ma124/strcase)
[![Godoc Reference](https://godoc.org/github.com/Ma124/strcase?status.svg)](http://godoc.org/github.com/Ma124/strcase)

strcase is a go package for converting string case to [snake case](https://en.wikipedia.org/wiki/Snake_case) or [camel case](https://en.wikipedia.org/wiki/CamelCase).

## Example

```go
x := "AnyKind of_string"
```

| Function                          | Result               |
|-----------------------------------|----------------------|
| `ToSnake(s)`                      | `any_kind_of_string` |
| `ToScreamingSnake(s)`             | `ANY_KIND_OF_STRING` |
| `ToKebab(s)`                      | `any-kind-of-string` |
| `ToScreamingKebab(s)`             | `ANY-KIND-OF-STRING` |
| `ToDelimited(s, '.')`             | `any.kind.of.string` |
| `ToScreamingDelimited(s, '.')`    | `ANY.KIND.OF.STRING` |
| `ToCamel(s)`                      | `AnyKindOfString`    |
| `ToLowerCamel(s)`                 | `anyKindOfString`    |
