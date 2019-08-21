package utils

import (
	"fmt"
	"os"
)

// DASH return dashes with fixed lenght
func DASH() string {
	return "--------------------------------------------------"
}

// Println prints text on the console
func Println(text string) {
	fmt.Println(text)
}

// PrintErr prints error on the console
func PrintErr(err error) {
	fmt.Println(err)
}

// ExitError exit the program with non success code
func ExitError() {
	os.Exit(1)
}

// SliceIndex find the index of element matching given predicate
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
