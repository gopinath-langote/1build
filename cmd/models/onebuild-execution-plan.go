package models

import sh "github.com/codeskyblue/go-sh"

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

func (executionPlan *OneBuildExecutionPlan) HasBefore() bool {
	return executionPlan.Before != nil
}

func (executionPlan *OneBuildExecutionPlan) HasAfter() bool {
	return executionPlan.After != nil
}

func (executionPlan *OneBuildExecutionPlan) HasCommands() bool {
	return len(executionPlan.Commands) > 0
}
