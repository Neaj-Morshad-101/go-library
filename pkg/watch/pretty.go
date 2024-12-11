package watch

import (
	"fmt"
	"reflect"
	"strings"
)

func prettyPrint(v interface{}) string {
	switch val := reflect.ValueOf(v); val.Kind() {
	case reflect.Struct:
		return prettyPrintStruct(val)
	case reflect.Map:
		return prettyPrintMap(val)
	case reflect.Slice, reflect.Array:
		return prettyPrintSlice(val)
	default:
		return fmt.Sprintf("%v", v)
	}
}

func prettyPrintStruct(v reflect.Value) string {
	t := v.Type()
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fields = append(fields, fmt.Sprintf("    %s: %v", field.Name, formatValue(value)))
	}
	return fmt.Sprintf("{\n%s\n}", strings.Join(fields, ",\n"))
}

func prettyPrintMap(v reflect.Value) string {
	var pairs []string
	for _, key := range v.MapKeys() {
		pairs = append(pairs, fmt.Sprintf("    %v: %v", key, formatValue(v.MapIndex(key))))
	}
	return fmt.Sprintf("{\n%s\n}", strings.Join(pairs, ",\n"))
}

func prettyPrintSlice(v reflect.Value) string {
	var elements []string
	for i := 0; i < v.Len(); i++ {
		elements = append(elements, formatValue(v.Index(i)))
	}
	return fmt.Sprintf("[\n    %s\n]", strings.Join(elements, " "))
}

func formatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		return prettyPrint(v.Interface())
	case reflect.Ptr:
		if v.IsNil() {
			return "nil"
		}
		return formatValue(v.Elem())
	case reflect.String:
		return fmt.Sprintf("%q", v.Interface())
	default:
		return fmt.Sprintf("%v", v.Interface())
	}
} 