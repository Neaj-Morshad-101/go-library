package watch

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

// captureOutput captures stdout during test execution
func captureOutput(f func()) string {
	// Save original stdout
	oldStdout := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function that generates output
	f()

	// Close the write end of the pipe
	w.Close()

	// Read the output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Restore original stdout
	os.Stdout = oldStdout

	return buf.String()
}

func TestWatch(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		contains []string // Strings that should be in the output
	}{
		{
			name:  "Integer",
			input: 42,
			contains: []string{
				"42",
				"watch_test.go",
			},
		},
		{
			name:  "String",
			input: "hello",
			contains: []string{
				`"hello"`,
				"watch_test.go",
			},
		},
		{
			name:  "Slice",
			input: []int{1, 2, 3},
			contains: []string{
				"Slice",
				"1",
				"2",
				"3",
			},
		},
		{
			name:  "Map",
			input: map[string]int{"a": 1, "b": 2},
			contains: []string{
				"Map",
				"a",
				"b",
				"1",
				"2",
			},
		},
		{
			name:  "Struct",
			input: struct{ Name string }{"test"},
			contains: []string{
				"Struct",
				`Name: "test"`,
			},
		},
		{
			name:     "Nil",
			input:    nil,
			contains: []string{"nil"},
		},
		{
			name:     "Pointer",
			input:    func() interface{} { x := 42; return &x }(),
			contains: []string{"Pointer", "42"},
		},
		{
			name:     "Nil Pointer",
			input:    (*int)(nil),
			contains: []string{"Pointer", "nil"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				Watch(tt.input)
			})

			// Strip color codes for easier testing
			output = stripColorCodes(output)

			// Check if output contains timestamp in format [HH:MM]
			if !strings.Contains(output, time.Now().Format("15:04")) {
				t.Errorf("Output missing timestamp: %s", output)
			}

			// Check if output contains expected strings
			for _, expected := range tt.contains {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', got: %s", expected, output)
				}
			}
		})
	}
}

// Add helper function to strip color codes
func stripColorCodes(s string) string {
	r := strings.NewReplacer(
		colorReset, "",
		colorRed, "",
		colorGreen, "",
		colorYellow, "",
		colorBlue, "",
		colorPurple, "",
		colorCyan, "",
		colorGray, "",
	)
	return r.Replace(s)
}

func TestExtractVariableName(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		line     int
		expected string
	}{
		{
			name:     "Invalid File",
			file:     "nonexistent.go",
			line:     1,
			expected: "<unknown>",
		},
		// Add more test cases if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractVariableName(tt.file, tt.line)
			if result != tt.expected {
				t.Errorf("extractVariableName() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Example test that will appear in the documentation
func ExampleWatch() {
	x := 42
	Watch(x)
	// Output will look similar to:
	// [12:34:56.789] (ExampleWatch) x = 42   // at watch_test.go:123
}

// Benchmark the Watch function
func BenchmarkWatch(b *testing.B) {
	x := 42
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Watch(x)
	}
}

// Test different types in a single struct
type testStruct struct {
	Int    int
	String string
	Map    map[string]int
	Slice  []string
}

func TestWatchComplexStruct(t *testing.T) {
	test := testStruct{
		Int:    42,
		String: "test",
		Map:    map[string]int{"a": 1},
		Slice:  []string{"one", "two"},
	}

	output := captureOutput(func() {
		Watch(test)
	})

	output = stripColorCodes(output)

	expectedFields := []string{
		"Int: 42",
		`String: "test"`,
		"Map: {",
		"a: 1",
		"Slice: [",
		`"one"`,
		`"two"`,
	}

	for _, field := range expectedFields {
		if !strings.Contains(output, field) {
			t.Errorf("Expected output to contain '%s', got: %s", field, output)
		}
	}
} 