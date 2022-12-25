# go-functional

### Functional Go library similar to Java's Stream & Optional API
All methods seem to work but I don't have the dedication to properly polish this project for actual use since it was just for fun.

For implemented methods see:
- [stream/stream.go](./stream/stream.go) for 'Stream'
- [check/check.go](./check/check.go) for 'Optional'

## TODO (but probably won't ever pick it up again)
- Make Stream's fully concurrent whilst still supporting short-circuit operations
- Unit-test

## Limitations
- Go's Generics don't allow for Generic input- and output- parameters on methods (not functions!)
-- Smooth chaining of methods such as Stream.map() in Java doesn't work
- Go's primitive types don't have null values and there is no equivalent "Object"
-- Optionals for boolean or int don't really make sense
