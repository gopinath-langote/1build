package models

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/codeskyblue/go-sh"
	"github.com/gopinath-langote/1build/cmd/utils"
)

// CommandContext holds all meta-data and required information for execution of a command.
type CommandContext struct {
	Name           string
	Command        string
	CommandSession *sh.Session
}

const (
	bannerOpen  = "[ "
	bannerClose = " ]"
)

// PrintPhaseBanner prints the CommandContext's name in a centred banner of the standard width.
// Similar-length command names will align when banners are printed sequentially.
func (c *CommandContext) PrintPhaseBanner() {
	centreLength := utf8.RuneCountInString(c.Name) +
		utf8.RuneCountInString(bannerOpen) +
		utf8.RuneCountInString(bannerClose)
	totalDashes := utils.MaxOutputWidth - centreLength

	// Intentional integer division — extra dash goes on the right so that
	// similar-length aliases line up neatly.
	numDashesLeft := totalDashes / 2
	numDashesRight := totalDashes / 2
	if totalDashes%2 == 1 {
		numDashesRight++
	}

	fmt.Print(strings.Repeat("-", numDashesLeft))
	fmt.Print(bannerOpen)
	utils.CPrint(c.Name, utils.Style{Color: utils.CYAN})
	fmt.Print(bannerClose)
	fmt.Print(strings.Repeat("-", numDashesRight))
	fmt.Println()
}
