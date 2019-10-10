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

	printResultsBanner(true, executeStart)
}

func executeAndStopIfFailed(command *models.CommandContext, executeStart time.Time) {
	command.PrintPhaseBanner()
	err := command.CommandSession.Run()
	if err != nil {
		exitCode := (err.Error())[12:]
		text := "\nExecution failed in phase '" + command.Name + "' – exit code: " + exitCode
		fmt.Println(utils.Colored(text, utils.RED))
		printResultsBanner(false, executeStart)
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
			utils.ExitWithCode("127")
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

// PrintResultsBanner prints result banner at the end of the test
func printResultsBanner(isSuccess bool, startTime time.Time) {
	timeDelta := time.Since(startTime)
	minutes := int64(timeDelta.Minutes())
	secs := int64(timeDelta.Seconds()) % 60
	var timeStr string
	if minutes == 0 {
		timeStr = fmt.Sprintf("%.2ds", secs)
	} else {
		timeStr = fmt.Sprintf("%.2dm %.2ds", minutes, secs)
	}
	result := utils.ColoredB("SUCCESS", utils.CYAN)
	if !isSuccess {
		result = utils.ColoredB("FAILURE", utils.RED)
	}
	result = fmt.Sprintf("%s - Total Time: %s", result, timeStr)
	fmt.Println("\n" + utils.Dash() + "\n" + result + "\n" + utils.Dash())
}
