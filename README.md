# Debug Watch Library

## Overview
This Go library provides a flexible `Watch()` function for easy debug logging of variables across multiple types.

## Installation

1. Create a new Go module (if not already created):
```bash
go mod init your-module-path
```

2. Create a `watch` directory in your project and save the `watch.go` file there.

## Usage

Import the package:
```go
import "your-module-path/watch"
```

Use the `Watch()` function to log variables:
```go
watch.Watch(yourVariable)
```

### Supported Types
- Primitive types (int, string, bool, etc.)
- Slices
- Maps
- Structs
- Pointers
- And more!

## Example Output
```
[main.go:10] main.main: 42
[main.go:13] main.main: Alice
[main.go:16] main.main: Slice (length: 5) [1 2 3 4 5]
```

## Features
- Automatically prints file name and line number
- Shows function context
- Handles various types dynamically
- Minimal performance overhead
```

## How to Use in Your Project

1. Copy the `watch.go` file into a `watch` directory in your project.
2. Ensure your `go.mod` file is set up correctly.
3. Import and use the `Watch()` function as shown in the example.

## Notes
- The library uses reflection for advanced type handling
- Best used for debugging and development
- Not recommended for production logging

## Customization
If you need to extend functionality, you can modify the `Watch()` function to add more type-specific handling.