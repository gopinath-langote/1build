package utils

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"strings"
)

// BANNER return dashes with fixed length - 72
func BANNER() string {
	return "------------------------------------------------------------------------"
}

// CPrintBold color print in bold
func CPrintBold(text string) {
	fmt.Println(aurora.BrightGreen(text).Bold())
}

// CPrintBoldUnderLine color print in bold and underline
func CPrintBoldUnderLine(text string) {
	aurora.BrightCyan(text).Bold().Underline()
}

// PrintErr prints error on the console
func PrintErr(err error) {
	fmt.Println(err)
}

// CPrintlnErr prints error line to console in bold Red
func CPrintlnErr(text string) {
	fmt.Println(aurora.Red("\n" + text).Bold())
}

// PrintlnDashedErr prints error line to console in bold Red with dashes above and below
func PrintlnDashedErr(text string) {
	errDash := strings.Repeat("-", len(text))
	fmt.Println()
	fmt.Println(errDash)
	fmt.Println(aurora.Red(text).Bold())
	fmt.Println(errDash)
}
