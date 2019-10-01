package exec

import (
	"fmt"
	"time"

	"github.com/codeskyblue/go-sh"
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/models"
	"github.com/gopinath-langote/1build/cmd/utils"
)

// ExecutePlan executes the Execution plan
func ExecutePlan(commands ...string) {

	executeStart := time.Now()

	configuration, err := config.LoadOneBuildConfiguration()
	if err != nil {
		fmt.Println(err)
		return
	}

	executionPlan := buildExecutionPlan(configuration, commands...)
	executionPlan.Print()

	if executionPlan.HasBefore() {
		executeAndStopIfFailed(executionPlan.Before, executeStart)
	}

	if executionPlan.HasCommands() {
		for _, commandContext := range executionPlan.Commands {
			executeAndStopIfFailed(commandContext, executeStart)
		}
	}

	if executionPlan.HasAfter() {
		executeAndStopIfFailed(executionPlan.After, executeStart)
	}

	utils.PrintResultsBanner(true, executeStart)
}

func executeAndStopIfFailed(command *models.CommandContext, executeStart time.Time) {
	command.PrintBanner()
	err := command.CommandSession.Run()
	if err != nil {
		exitCode := (err.Error())[12:]
		utils.PrintlnErr("Execution failed during phase \"" +
			command.Name +
			"\" - Execution of the script \"" +
			command.Command +
			"\" returned non-zero exit code : " + exitCode)
		utils.PrintResultsBanner(false, executeStart)
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
			fmt.Println(utils.ColoredB("\nError building execution plan. Command \""+name+"\" not found.", utils.RED))
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
