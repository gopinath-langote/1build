package utils

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"strings"
)

// BANNER return dashes with fixed length - 72
func BANNER() string {
	return strings.Repeat("-", MaxOutputWidth)
}

// OneBuildColor represents type for color enum
type OneBuildColor int

const (
	// CYAN is 1build's default color standard
	CYAN OneBuildColor = 0

	// RED is used in failure messages
	RED  OneBuildColor = 1
)

// ColoredB return text in color with bold format
func ColoredB(text string, color OneBuildColor) string {
	return colorize(text, color).Bold().String()
}

// Colored return text in color
func Colored(text string, color OneBuildColor) string {
	return colorize(text, color).String()
}

// ColoredBU return text in color with bold and underline format
func ColoredBU(text string, color OneBuildColor) string {
	return colorize(text, color).Bold().Underline().String()
}

// PrintlnDashedErr prints error line to console in bold Red with dashes above and below
func PrintlnDashedErr(text string) {
	errDash := strings.Repeat("-", len(text))
	fmt.Println()
	fmt.Println(errDash)
	fmt.Println(ColoredB(text, RED))
	fmt.Println(errDash)
}

func colorize(text string, color OneBuildColor) aurora.Value {
	var coloredText aurora.Value
	if color == CYAN {
		coloredText = aurora.BrightCyan(text)
	} else {
		coloredText = aurora.Red(text)
	}
	return coloredText
}
