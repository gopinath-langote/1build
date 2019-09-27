package utils

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/gopinath-langote/1build/cmd/models"
	"github.com/logrusorgru/aurora"
	rune "github.com/mattn/go-runewidth"
)

// DASH return dashes with fixed length
func DASH() string {
	return "--------------------------------------------------"
}

// DashesMatchingTextLength returns a string of '-' matching the length of provided string
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
	fmt.Println(err)
}

// PrintlnErr prints error line to console in bold Red
func PrintlnErr(text string) {
	fmt.Println(aurora.Red("\n" + text).Bold())
}

// PrintlnDashedErr prints error line to console in bold Red with dashes above and below
func PrintlnDashedErr(text string) {
	errDash := DashesMatchingTextLength(text)
	fmt.Println()
	fmt.Println(errDash)
	fmt.Println(aurora.Red(text).Bold())
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
	fmt.Println(aurora.BrightGreen("Execution plan (executed in ordered sequence)").Bold().Underline())
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
	fmt.Println()
	fmt.Println()
}
