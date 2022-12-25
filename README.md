# go-functional

### Functional Go library similar to Java's Stream & Optional API
All methods seem to work but I don't have the dedication to properly polish this project for actual use since it was just for fun.

For implemented methods see:
* [stream/stream.go](./stream/stream.go) for "Stream"
* [check/check.go](./check/check.go) for "Optional"

## TODO (but probably won't ever pick it up again)
* Make Stream's fully concurrent whilst still supporting short-circuit operations
* Unit-test

## Limitations
* Go's Generics don't allow for generic input- and output- types on methods (not functions!)
  * Methods such as Stream.map() in Java can't be smoothly chained or must have a predefined output type
* Go's primitive types don't have null values and there is no equivalent "Object"
  * Optionals for types like boolean or int don't really make sense
