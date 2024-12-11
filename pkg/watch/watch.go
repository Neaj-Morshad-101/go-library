package watch

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// Watch is a generic debug logging function that prints variable details
func Watch(v interface{}) {
	// Get timestamp
	timestamp := time.Now().Format("15:04:05.000")

	// Get caller information
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Unable to get caller information")
		return
	}

	// Extract function name and clean it
	fn := runtime.FuncForPC(pc)
	funcName := "unknown"
	if fn != nil {
		fullName := fn.Name()
		if lastDot := strings.LastIndex(fullName, "."); lastDot != -1 {
			funcName = fullName[lastDot+1:]
		} else {
			funcName = fullName
		}
	}

	// Get filename
	shortFile := file
	if lastSlashIndex := strings.LastIndexByte(file, '/'); lastSlashIndex != -1 {
		shortFile = file[lastSlashIndex+1:]
	}

	// Get variable name
	varName := extractVariableName(file, line)

	// Format the output with colors
	timestampStr := colorize("["+timestamp+"]", colorBlue)
	funcNameStr := colorize("("+funcName+")", colorYellow)
	varNameStr := colorize(varName, colorGreen)
	locationStr := colorize(fmt.Sprintf("// at %s:%d", shortFile, line), colorGray)

	// Print based on type with colors
	switch val := v.(type) {
	case nil:
		fmt.Printf("%s %s %s = %s   %s\n",
			timestampStr, funcNameStr, varNameStr,
			colorize("nil", colorRed), locationStr)
	case bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128:
		fmt.Printf("%s %s %s = %s   %s\n",
			timestampStr, funcNameStr, varNameStr,
			colorize(fmt.Sprintf("%v", val), colorCyan), locationStr)
	case string:
		fmt.Printf("%s %s %s = %s   %s\n",
			timestampStr, funcNameStr, varNameStr,
			colorize(fmt.Sprintf("%q", val), colorCyan), locationStr)
	default:
		v := reflect.ValueOf(v)
		switch v.Kind() {
		case reflect.Slice:
			fmt.Printf("%s %s %s = %s %s   %s\n",
				timestampStr, funcNameStr, varNameStr,
				colorize("Slice", colorPurple),
				colorize(prettyPrint(v.Interface()), colorCyan),
				locationStr)
		case reflect.Map:
			fmt.Printf("%s %s %s = %s %s   %s\n",
				timestampStr, funcNameStr, varNameStr,
				colorize("Map", colorPurple),
				colorize(prettyPrint(v.Interface()), colorCyan),
				locationStr)
		case reflect.Struct:
			fmt.Printf("%s %s %s = %s %s   %s\n",
				timestampStr, funcNameStr, varNameStr,
				colorize("Struct", colorPurple),
				colorize(prettyPrint(v.Interface()), colorCyan),
				locationStr)
		case reflect.Ptr:
			if v.IsNil() {
				fmt.Printf("%s %s %s = %s   %s\n",
					timestampStr, funcNameStr, varNameStr,
					colorize("Pointer (nil)", colorRed), locationStr)
			} else {
				fmt.Printf("%s %s %s = %s %s   %s\n",
					timestampStr, funcNameStr, varNameStr,
					colorize("Pointer", colorPurple),
					colorize(prettyPrint(v.Elem().Interface()), colorCyan),
					locationStr)
			}
		default:
			fmt.Printf("%s %s %s = %s   %s\n",
				timestampStr, funcNameStr, varNameStr,
				colorize(fmt.Sprintf("%v", v.Interface()), colorCyan),
				locationStr)
		}
	}
}

// extractVariableName attempts to extract the variable name from the source code
func extractVariableName(file string, line int) string {
	// Read the source file
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "<unknown>"
	}

	// Parse the source file
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return "<unknown>"
	}

	// Find the Watch call at the specified line
	var varName string
	ast.Inspect(f, func(n ast.Node) bool {
		if call, ok := n.(*ast.CallExpr); ok {
			if fset.Position(call.Pos()).Line == line {
				if len(call.Args) > 0 {
					if ident, ok := call.Args[0].(*ast.Ident); ok {
						varName = ident.Name
					}
				}
			}
		}
		return true
	})

	if varName == "" {
		return "<unknown>"
	}
	return varName
}
