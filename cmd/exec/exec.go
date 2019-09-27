package exec

import (
	"fmt"

	"github.com/codeskyblue/go-sh"
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/models"
	"github.com/gopinath-langote/1build/cmd/utils"
)

// ExecutePlan executes the Execution plan
func ExecutePlan(commands ...string) {

	configuration, err := config.LoadOneBuildConfiguration()
	if err != nil {
		utils.PrintErr(err)
		return
	}

	executionPlan := buildExecutionPlan(configuration, commands...)
	executionPlan.Print()

	if executionPlan.HasBefore() {
		executeAndStopIfFailed(executionPlan.Before)
	}

	if executionPlan.HasCommands() {
		for _, commandContext := range executionPlan.Commands {
			executeAndStopIfFailed(commandContext)
		}
	}

	if executionPlan.HasAfter() {
		executeAndStopIfFailed(executionPlan.After)
	}

	fmt.Println()
	utils.CPrintBold("SUCCESS")

}

func executeAndStopIfFailed(command *models.CommandContext) {
	command.PrintBanner()
	err := command.CommandSession.Run()
	if err != nil {
		exitCode := (err.Error())[12:]
		utils.PrintlnDashedErr("Execution failed during phase \"" + command.Name + "\" - Execution of the script \"" + command.Command + "\" returned non-zero exit code : " + exitCode)
		utils.ExitWithCode(exitCode)
	}
}

func buildExecutionPlan(onebuildConfig config.OneBuildConfiguration, commands ...string) models.OneBuildExecutionPlan {

	before := onebuildConfig.Before
	var executionPlan models.OneBuildExecutionPlan
	if before != "" {
		executionPlan.Before = &models.CommandContext{
			Name: "before", Command: before, CommandSession: bashCommand(sh.NewSession(), before)}
	}

	for _, name := range commands {
		executionCommand := onebuildConfig.GetCommand(name)
		if executionCommand == "" {
			utils.CPrintlnErr("Error building execution plan. Command \"" + name + "\" not found.")
			onebuildConfig.Print()
			utils.Exit(127)
		}
		executionPlan.Commands = append(executionPlan.Commands, &models.CommandContext{
			Name: name, Command: executionCommand, CommandSession: bashCommand(sh.NewSession(), executionCommand)})
	}

	after := onebuildConfig.After
	if after != "" {
		executionPlan.After = &models.CommandContext{
			Name: "after", Command: after, CommandSession: bashCommand(sh.NewSession(), after)}
	}

	return executionPlan
}

func bashCommand(s *sh.Session, command string) *sh.Session {
	return s.Command("bash", "-c", command)
}
