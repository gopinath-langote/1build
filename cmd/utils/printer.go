package utils

import (
	"strings"

	"github.com/fatih/color"
)

// Dash return dashes with fixed length - 72
func Dash() string {
	return strings.Repeat("-", MaxOutputWidth)
}

const (
	// CYAN is 1build's default color standard
	CYAN color.Attribute = color.FgCyan

	// RED is used in failure messages
	RED color.Attribute = color.FgHiRed
)

// Style represents type for text formatting
type Style struct {
	Color color.Attribute
	Bold  bool
}

// CPrintln prints the text with given formatting style with newline
func CPrintln(text string, style Style) {
	format(style).Println(text)
}

// CPrint prints the text with given formatting style
func CPrint(text string, style Style) {
	format(style).Print(text)
}

func format(style Style) *color.Color {
	formatter := color.New(style.Color)
	if style.Bold {
		formatter = formatter.Add(color.Bold)
	}
	return formatter
}
