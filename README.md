strcase [![Build Status](https://travis-ci.org/iancoleman/strcase.svg)](https://travis-ci.org/iancoleman/strcase) [![Coverage](http://gocover.io/_badge/github.com/iancoleman/strcase?0)](http://gocover.io/github.com/iancoleman/strcase)
=======

strcase is a go package for converting string case to [snake case](https://en.wikipedia.org/wiki/Snake_case) or [camel case](https://en.wikipedia.org/wiki/CamelCase).

# Example

```
x := "AnyKind of_string"
xSnake := strcase.ToSnake(x) // any_kind_of_string
xCamel := strcase.ToCamel(x) // AnyKindOfString
```
