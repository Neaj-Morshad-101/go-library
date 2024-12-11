package watch

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// Watch is a generic debug logging function that prints variable details
func Watch(v interface{}) {
	// Get caller information
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Unable to get caller information")
		return
	}

	// Extract function name
	fn := runtime.FuncForPC(pc)
	funcName := "unknown"
	if fn != nil {
		funcName = fn.Name()
	}

	// Extract filename
	shortFile := file
	if lastSlashIndex := strings.LastIndexByte(file, '/'); lastSlashIndex != -1 {
		shortFile = file[lastSlashIndex+1:]
	}

	// Print based on type
	switch val := v.(type) {
	case nil:
		fmt.Printf("[%s:%d] %s: nil\n", shortFile, line, funcName)
	case bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, float32, float64, complex64, complex128, string:
		fmt.Printf("[%s:%d] %s: %v\n", shortFile, line, funcName, val)
	case []interface{}:
		fmt.Printf("[%s:%d] %s: Slice %v (length: %d)\n", shortFile, line, funcName, val, len(val))
	case map[interface{}]interface{}:
		fmt.Printf("[%s:%d] %s: Map %v (length: %d)\n", shortFile, line, funcName, val, len(val))
	case struct{}:
		fmt.Printf("[%s:%d] %s: Struct %+v\n", shortFile, line, funcName, val)
	default:
		// Handle more complex types using reflection
		v := reflect.ValueOf(v)
		switch v.Kind() {
		case reflect.Slice:
			fmt.Printf("[%s:%d] %s: Slice (length: %d) %v\n",
				shortFile, line, funcName, v.Len(), v.Interface())
		case reflect.Map:
			fmt.Printf("[%s:%d] %s: Map (length: %d) %v\n",
				shortFile, line, funcName, v.Len(), v.Interface())
		case reflect.Struct:
			fmt.Printf("[%s:%d] %s: Struct %+v\n",
				shortFile, line, funcName, v.Interface())
		case reflect.Ptr:
			if v.IsNil() {
				fmt.Printf("[%s:%d] %s: Pointer (nil)\n", shortFile, line, funcName)
			} else {
				fmt.Printf("[%s:%d] %s: Pointer %v\n",
					shortFile, line, funcName, v.Elem().Interface())
			}
		default:
			fmt.Printf("[%s:%d] %s: %v (Type: %v)\n",
				shortFile, line, funcName, v.Interface(), v.Type())
		}
	}
}
