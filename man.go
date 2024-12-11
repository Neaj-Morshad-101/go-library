package go_library

import (
	"fmt"
	"github.com/Neaj-Morshad-101/go-library/pkg/watch"
	"math/rand"
	"time"
)

//package main

// ComplexStruct to demonstrate struct watching
type ComplexStruct struct {
	ID       int
	Name     string
	Active   bool
	Metadata map[string]string
}

// processData simulates a function with multiple variables to watch
func processData(input []int) []int {
	watch.Watch(input) // Watch input slice

	// Some data processing
	var processed []int
	for _, num := range input {
		if num%2 == 0 {
			processed = append(processed, num*2)
		}
	}

	watch.Watch(processed) // Watch processed slice
	return processed
}

// generateRandomData creates test data
func generateRandomData() *ComplexStruct {
	rand.Seed(time.Now().UnixNano())

	return &ComplexStruct{
		ID:     rand.Intn(1000),
		Name:   fmt.Sprintf("User-%d", rand.Intn(100)),
		Active: rand.Intn(2) == 1,
		Metadata: map[string]string{
			"created_at": time.Now().Format(time.RFC3339),
			"source":     "random_generator",
		},
	}
}

func main() {
	// Primitive type watching
	age := 30
	watch.Watch(age)

	name := "John Doe"
	watch.Watch(name)

	// Slice watching
	numbers := []int{1, 2, 3, 4, 5}
	watch.Watch(numbers)

	// Map watching
	scores := map[string]int{
		"math":    95,
		"science": 88,
		"history": 92,
	}
	watch.Watch(scores)

	// Pointer watching
	var ptr *int
	watch.Watch(ptr)

	ptr = &age
	watch.Watch(ptr)

	// Struct watching
	person := ComplexStruct{
		ID:       1,
		Name:     "Alice",
		Active:   true,
		Metadata: map[string]string{"role": "admin"},
	}
	watch.Watch(person)

	// Function with internal variable watching
	processedNumbers := processData([]int{2, 3, 4, 5, 6})
	watch.Watch(processedNumbers)

	// Random data generation and watching
	randomUser := generateRandomData()
	watch.Watch(randomUser)

	// Nil type watching
	var emptySlice []string
	watch.Watch(emptySlice)

	var nilMap map[string]int
	watch.Watch(nilMap)
}
