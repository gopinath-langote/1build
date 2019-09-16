package exec

import (
	"fmt"

	sh "github.com/codeskyblue/go-sh"
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/models"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/logrusorgru/aurora"
)

// ExecutePlan executes the Execution plan
func ExecutePlan(commands ...string) {

	configuration, err := config.LoadOneBuildConfiguration()
	if err != nil {
		utils.PrintErr(err)
		return
	}

	executionPlan := buildExecutionPlan(configuration, commands...)
	utils.PrintExecutionPlan(executionPlan)

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
	fmt.Println(aurora.BrightGreen("SUCCESS").Bold())

}

func getCommandFromName(config config.OneBuildConfiguration, cmd string) string {
	for _, command := range config.Commands {
		for k, v := range command {
			if k == cmd {
				return v
			}
		}
	}
	return ""
}

func bashCommand(s *sh.Session, command string) *sh.Session {
	return s.Command("bash", "-c", command)
}

func executeAndStopIfFailed(command *models.CommandContext) {
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
		executionPlan.Before = &models.CommandContext{"before", before, bashCommand(sh.NewSession(), before)}
	}

	for _, name := range commands {
		executionCommand := getCommandFromName(onebuildConfig, name)
		if executionCommand == "" {
			utils.PrintlnErr("Error building exectuion plan. Command \"" + name + "\" not found.")
			config.PrintConfiguration(onebuildConfig)
			utils.Exit(1)
		}
		executionPlan.Commands = append(executionPlan.Commands, &models.CommandContext{name, executionCommand, bashCommand(sh.NewSession(), executionCommand)})
	}

	after := onebuildConfig.After
	if after != "" {
		executionPlan.After = &models.CommandContext{"after", after, bashCommand(sh.NewSession(), after)}
	}

	return executionPlan
}
