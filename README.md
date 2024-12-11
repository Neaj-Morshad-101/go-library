# Go Debug Watch

A lightweight, colorful debugging library for Go that provides pretty-printed variable inspection during development.

## Features

- 🎨 Colorized output for better readability
- 🔍 Automatic variable name detection
- ⏰ Timestamp for each watch call
- 📍 File and line number tracking
- 🎯 Function name context
- 🎭 Pretty printing for complex types
- 💪 Type-safe with Go generics
- 🚀 Zero external dependencies

## Installation

```bash
go get github.com/Neaj-Morshad-101/go-library/pkg/watch
```

## Usage

```go
import "github.com/Neaj-Morshad-101/go-library/pkg/watch"

func main() {
    // Watch primitive types
    name := "John Doe"
    watch.Watch(name)
    // Output: [12:34:56.789] (main) name = "John Doe"   // at main.go:6

    // Watch slices
    numbers := []int{1, 2, 3}
    watch.Watch(numbers)
    // Output: [12:34:56.789] (main) numbers = Slice [
    //     1 2 3
    // ]   // at main.go:10

    // Watch maps
    scores := map[string]int{"math": 95, "science": 88}
    watch.Watch(scores)
    // Output: [12:34:56.789] (main) scores = Map {
    //     math: 95,
    //     science: 88
    // }   // at main.go:14

    // Watch structs
    type Person struct {
        Name string
        Age  int
    }
    person := Person{Name: "Alice", Age: 30}
    watch.Watch(person)
    // Output: [12:34:56.789] (main) person = Struct {
    //     Name: "Alice",
    //     Age: 30
    // }   // at main.go:23
}
```

## Output Format

Each watch call produces a line of output with the following components:

- Timestamp: `[HH:MM:SS.mmm]`
- Function name: `(functionName)`
- Variable name: Automatically detected from source
- Value: Pretty-printed with type information
- Location: File and line number

## Color Scheme

- 🔵 Blue: Timestamps
- 🟡 Yellow: Function names
- 🟢 Green: Variable names
- 🔴 Red: Nil values
- 🟣 Purple: Type information
- 🟦 Cyan: Values
- ⬜ Gray: File locations

## Supported Types

- All primitive types
- Strings
- Slices and Arrays
- Maps
- Structs
- Pointers
- Complex types (nested structures)
- nil values

## Development

### Running Tests

```bash
go test ./pkg/watch
```

### Benchmarks

```bash
go test -bench=. ./pkg/watch
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Acknowledgments

- Inspired by the debug patterns in various programming languages
- Built with ❤️ for the Go community