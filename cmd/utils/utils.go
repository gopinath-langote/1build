package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	rune "github.com/mattn/go-runewidth"
)

var red = color.New(color.FgRed)
var boldRed = red.Add(color.Bold)

// DASH return dashes with fixed lenght
func DASH() string {
	return "--------------------------------------------------"
}

// DashesMatchingTextLength returns a string of '-' matching the length of provided stirng
func DashesMatchingTextLength(text string) string {
	width := rune.StringWidth(text)
	dashString := ""
	for i := 1; i <= width; i++ {
		dashString = dashString + "-"
	}
	return dashString
}

// Println prints text on the console
func Println(text string) {
	fmt.Println(text)
}

// PrintErr prints error on the console
func PrintErr(err error) {
	boldRed.Println(err)
}

// PrintlnErr prints error line to console in bold Red
func PrintlnErr(text string) {
	boldRed.Println("\n" + text)
}

// PrintlnDashedErr prints error line to console in bold Red with dashes above and below
func PrintlnDashedErr(text string) {
	errDash := DashesMatchingTextLength(text)
	fmt.Println()
	fmt.Println(errDash)
	boldRed.Println(text)
	fmt.Println(errDash)
}

// Exit exits the current process with specified exit code
func Exit(code int) {
	os.Exit(code)
}

// ExitWithCode exits the current process with specified exit code provided as stirng
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
