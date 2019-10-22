package utils

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
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

// Style represents type for text formatting
type Style struct {
	Color OneBuildColor
	Bold  bool
}

// CPrintln prints the text with given formatting style with newline
func CPrintln(text string, style Style) {
	CPrint(text, style)
	fmt.Println()
}

// CPrint prints the text with given formatting style
func CPrint(text string, style Style) {
	formattedText := colorize(text, style)
	formattedText = bold(formattedText, style)
	fmt.Print(formattedText.String())
}

func bold(formattedText aurora.Value, style Style) aurora.Value {
	if style.Bold {
		return formattedText.Bold()
	}
	return formattedText
}

func colorize(text string, style Style) aurora.Value {
	var coloredText aurora.Value
	if style.Color == CYAN {
		coloredText = aurora.BrightCyan(text)
	} else {
		coloredText = aurora.BrightRed(text)
	}
	return coloredText
}
