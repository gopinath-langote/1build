package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/logrusorgru/aurora"
	"golang.org/x/term"
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

// colorEnabled returns true when color output should be applied.
// Color is disabled when the NO_COLOR env var is set (https://no-color.org/)
// or when stdout is not connected to a terminal.
func colorEnabled() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// CPrintln prints the text with given formatting style with newline to stdout.
func CPrintln(text string, style Style) {
	CPrint(text, style)
	fmt.Println()
}

// CPrint prints the text with given formatting style to stdout.
func CPrint(text string, style Style) {
	fmt.Print(format(text, style))
}

// CPrintlnErr prints the text with given formatting style with newline to stderr.
func CPrintlnErr(text string, style Style) {
	CPrintErr(text, style)
	fmt.Fprintln(os.Stderr)
}

// CPrintErr prints the text with given formatting style to stderr.
func CPrintErr(text string, style Style) {
	fmt.Fprint(os.Stderr, format(text, style))
}

// format applies color and bold styling to text, respecting NO_COLOR / TTY.
func format(text string, style Style) string {
	if !colorEnabled() {
		return text
	}
	v := colorize(text, style)
	v = bold(v, style)
	return v.String()
}

func bold(formattedText aurora.Value, style Style) aurora.Value {
	if style.Bold {
		return formattedText.Bold()
	}
	return formattedText
}

func colorize(text string, style Style) aurora.Value {
	if style.Color == CYAN {
		return aurora.BrightCyan(text)
	}
	return aurora.BrightRed(text)
}
