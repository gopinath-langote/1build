package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/gopinath-langote/1build/cmd/models"
	"github.com/logrusorgru/aurora"
	runeWidth "github.com/mattn/go-runewidth"
)

var colors aurora.Aurora

func init() {
	enableColors, _ := os.LookupEnv("DISABLE_1B_COLORS")
	colors = aurora.NewAurora(enableColors == "" || enableColors == "true")
}

// DASH return dashes with fixed length
func DASH() string {
	return "--------------------------------------------------"
}

// DashesMatchingTextLength returns a string of '-' matching the length of provided string
func DashesMatchingTextLength(text string) string {
	width := runeWidth.StringWidth(text)
	return dashOfLength(width)
}

func dashOfLength(length int) string {
	dashString := make([]byte, length)
	for i := 0; i < length; i++ {
		dashString[i] = '-'
	}
	return string(dashString)
}

// Println prints text on the console
func Println(text string) {
	fmt.Println(text)
}

// PrintErr prints error on the console
func PrintErr(err error) {
	fmt.Println(err)
}

//PrintSuccessBold prints success message (in bright greed bold)
func PrintSuccessBold(s string) {
	fmt.Println(colors.BrightGreen(s).Bold())
}

// PrintlnErr prints error line to console in bold Red
func PrintlnErr(text string) {
	fmt.Println(colors.Red("\n" + text).Bold())
}

// PrintlnDashedErr prints error line to console in bold Red with dashes above and below
func PrintlnDashedErr(text string) {
	errDash := DashesMatchingTextLength(text)
	fmt.Println()
	fmt.Println(errDash)
	fmt.Println(colors.Red(text).Bold())
	fmt.Println(errDash)
}

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

// PrintExecutionPlan prints a formatted execution plan to console
func PrintExecutionPlan(executionPlan models.OneBuildExecutionPlan) {
	fmt.Println()
	fmt.Println(colors.BrightGreen("Execution plan (executed in ordered sequence)").Bold().Underline())
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.TabIndent)

	maxPhaseName := "Phase"
	maxCommand := "Command"

	if executionPlan.HasBefore() {
		maxPhaseName = executionPlan.Before.Name
		maxCommand = executionPlan.Before.Command
	}

	if executionPlan.HasCommands() {
		for _, command := range executionPlan.Commands {
			if len(command.Name) > len(maxPhaseName) {
				maxPhaseName = command.Name
			}
			if len(command.Command) > len(maxCommand) {
				maxCommand = command.Command
			}
		}
	}

	if executionPlan.HasAfter() {
		command := executionPlan.After
		if len(command.Name) > len(maxPhaseName) {
			maxPhaseName = command.Name
		}
		if len(command.Command) > len(maxCommand) {
			maxCommand = command.Command
		}
	}

	phaseDashes := DashesMatchingTextLength(maxPhaseName)
	commandDashes := DashesMatchingTextLength(maxCommand)

	fmt.Fprintf(w, "%s\t%s\n", phaseDashes, commandDashes)
	fmt.Fprintln(w, "Phase\tCommand")
	fmt.Fprintf(w, "%s\t%s\n", phaseDashes, commandDashes)

	if executionPlan.HasBefore() {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s", executionPlan.Before.Name, executionPlan.Before.Command))
	}

	if executionPlan.HasCommands() {
		for _, command := range executionPlan.Commands {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s", command.Name, command.Command))
		}
	}

	if executionPlan.HasAfter() {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s", executionPlan.After.Name, executionPlan.After.Command))
	}

	w.Flush()
}

//PrintPhaseBanner prints banner for phase
func PrintPhaseBanner(command *models.CommandContext) {
	coloredPhaseName := colors.Cyan("[ " + command.Name + " ]").String()
	coloredPhaseNameWidth := runeWidth.StringWidth(coloredPhaseName)
	dashPrefixes := dashOfLength((72 - coloredPhaseNameWidth) / 2)
	dashPostfixes := dashOfLength(72 - runeWidth.StringWidth(dashPrefixes+coloredPhaseName))
	finalBanner := fmt.Sprintf("\n%s%s%s", dashPostfixes, coloredPhaseName, dashPostfixes)
	if runeWidth.StringWidth(finalBanner) != 72 {
		finalBanner = strings.TrimSuffix(finalBanner, "-")
	}
	fmt.Println(finalBanner)
}
