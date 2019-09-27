package utils

import (
	"os"
	"strconv"
)

// Exit exits the current process with specified exit code
func Exit(code int) {
	os.Exit(code)
}

// ExitWithCode exits the current process with specified exit code provided as string
func ExitWithCode(code string) {
	exitCode, _ := strconv.Atoi(code)
	os.Exit(exitCode)
}

// ExitError exit the program with non success code
func ExitError() {
	Exit(1)
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

// LongestString return longest string from array
func LongestString(list []string) (longest string) {
	for _, item := range list {
		if len(item) > len(longest) {
			longest = item
		}
	}
	return longest
}
