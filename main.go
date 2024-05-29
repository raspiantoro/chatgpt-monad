package main

import (
	"fmt"
)

// Monad interface with type parameters T and U
type Monad[T any, U any] interface {
	Return(value T) Monad[T, U]
	Bind(func(T) Monad[U, U]) Monad[U, U]
}

// Maybe monad with type parameters T and U
type Maybe[T any, U any] struct {
	value T    // Value of type T
	empty bool // Indicates if value is uninitialized
}

// Return lifts a value into the Maybe monad
func (m Maybe[T, U]) Return(value T) Monad[T, U] {
	return Maybe[T, U]{value: value, empty: false}
}

// Bind applies a function to the value inside the Maybe monad
func (m Maybe[T, U]) Bind(f func(T) Monad[U, U]) Monad[U, U] {
	if m.empty {
		return Maybe[U, U]{empty: true}
	}
	return f(m.value)
}

// Example usage
func main() {
	// Example 1: Using Maybe monad with integer type
	var monad1 Monad[int, string] = Maybe[int, string]{empty: true}
	monad1 = monad1.Return(42)
	var monad2 Monad[string, string] = monad1.Bind(func(value int) Monad[string, string] {
		// Perform a calculation before converting to string
		calculatedValue := value * 2
		strValue := fmt.Sprintf("%d", calculatedValue)
		return Maybe[string, string]{value: strValue, empty: false}
	})
	fmt.Println(monad2) // Output: {84}

	// Example 2: Using Maybe monad with string type
	var monad3 Monad[string, int] = Maybe[string, int]{empty: true}
	monad3 = monad3.Return("Hello")
	var monad4 Monad[int, int] = monad3.Bind(func(value string) Monad[int, int] {
		length := len(value)
		return Maybe[int, int]{value: length, empty: false}
	})
	fmt.Println(monad4) // Output: {5}
}
