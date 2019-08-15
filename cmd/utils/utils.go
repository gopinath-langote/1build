package utils

import (
	"fmt"
	"os"
)

func DASH() string {
	return "--------------------------------------------------"
}

func Println(text string) {
	fmt.Println(text)
}

func PrintErr(err error) {
	fmt.Println(err)
}

func ExitError() {
	os.Exit(1)
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}
