package utils

import (
	"github.com/logrusorgru/aurora"
	"strings"
)

// Dash return dashes with fixed length - 72
func Dash() string {
	return strings.Repeat("-", MaxOutputWidth)
}

// OneBuildColor represents type for color enum
type OneBuildColor int

const (
	// CYAN is 1build's default color standard
	CYAN OneBuildColor = 0

	// RED is used in failure messages
	RED OneBuildColor = 1
)

// ColoredB return text in color with bold format
func ColoredB(text string, color OneBuildColor) string {
	return colorize(text, color).Bold().String()
}

// Colored return text in color
func Colored(text string, color OneBuildColor) string {
	return colorize(text, color).String()
}

// ColoredU return text in color with bold and underline format
func ColoredU(text string, color OneBuildColor) string {
	return colorize(text, color).Underline().String()
}

func colorize(text string, color OneBuildColor) aurora.Value {
	var coloredText aurora.Value
	if color == CYAN {
		coloredText = aurora.BrightCyan(text)
	} else {
		coloredText = aurora.BrightRed(text)
	}
	return coloredText
}
