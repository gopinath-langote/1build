package models

import (
	"fmt"
	"github.com/codeskyblue/go-sh"
	"github.com/gopinath-langote/1build/cmd/utils"
	"os"
	"strings"
	"text/tabwriter"
)

// OneBuildExecutionPlan holds all information for the execution strategy
type OneBuildExecutionPlan struct {
	Before   *CommandContext
	Commands []*CommandContext
	After    *CommandContext
}

// CommandContext holds all meta-data and required information for execution of a command
type CommandContext struct {
	Name           string
	Command        string
	CommandSession *sh.Session
}

// HasBefore return true if plan contains before section else false
func (executionPlan *OneBuildExecutionPlan) HasBefore() bool {
	return executionPlan.Before != nil
}

// HasAfter return true if plan contains after section else false
func (executionPlan *OneBuildExecutionPlan) HasAfter() bool {
	return executionPlan.After != nil
}

// HasCommands return true if plan contains command(s) else false
func (executionPlan *OneBuildExecutionPlan) HasCommands() bool {
	return len(executionPlan.Commands) > 0
}

// Print prints execution plan
func (executionPlan *OneBuildExecutionPlan) Print() {
	fmt.Println()
	utils.CPrintBoldUnderLine("Execution plan")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.TabIndent)

	phase, cmd := longestPhaseAndCommandValue(executionPlan)
	fmt.Fprintf(w, "%s\t%s\n", dashesOfLength(phase), dashesOfLength(cmd))
	fmt.Fprintln(w, "Phase\tCommand")
	fmt.Fprintf(w, "%s\t%s\n", dashesOfLength(phase), dashesOfLength(cmd))

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

	_ = w.Flush()
	fmt.Print("\n\n")
}

func longestPhaseAndCommandValue(executionPlan *OneBuildExecutionPlan) (string, string) {
	var phases []string
	var cmdValues []string
	if executionPlan.HasBefore() {
		phases = append(phases, executionPlan.Before.Name)
		cmdValues = append(cmdValues, executionPlan.Before.Command)
	}
	if executionPlan.HasCommands() {
		for _, command := range executionPlan.Commands {
			phases = append(phases, command.Name)
			cmdValues = append(cmdValues, command.Command)
		}
	}
	if executionPlan.HasAfter() {
		phases = append(phases, executionPlan.After.Name)
		cmdValues = append(cmdValues, executionPlan.After.Command)
	}
	return utils.LongestString(phases), utils.LongestString(cmdValues)
}

func dashesOfLength(text string) string {
	return strings.Repeat("-", len(text))
}
