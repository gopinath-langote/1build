package exec

import (
	"fmt"
	"os"
	"text/tabwriter"

	sh "github.com/codeskyblue/go-sh"
	"github.com/fatih/color"
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
)

var boldGreeen = color.New(color.FgHiGreen).Add(color.Bold)

// ExecutePlan executes the Execution plan
func ExecutePlan(commands ...string) {

	configuration, err := config.LoadOneBuildConfiguration()
	if err != nil {
		utils.PrintErr(err)
		return
	}

	executionPlan := buildExecutionPlan(configuration, commands...)
	executionPlan.print()

	if executionPlan.hasInit() {
		executeAndStopIfFailed(executionPlan.Init)
	}

	if executionPlan.hasCommands() {
		for _, commandContext := range executionPlan.Commands {
			executeAndStopIfFailed(commandContext)
		}
	}

	if executionPlan.hasDestroy() {
		executeAndStopIfFailed(executionPlan.Destroy)
	}

	fmt.Println()
	boldGreeen.Println("SUCCESS")

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

func executeAndStopIfFailed(command *CommandContext) {
	err := command.commandSession.Run()
	if err != nil {
		exitCode := (err.Error())[12:]
		utils.PrintlnDashedErr("Execution failed during phase \"" + command.name + "\" - Execution of the script \"" + command.command + "\" returned non-zero exit code : " + exitCode)
		utils.ExitWithCode(exitCode)
	}
}

func buildExecutionPlan(onebuildConfig config.OneBuildConfiguration, commands ...string) OneBuildExecutionPlan {

	before := onebuildConfig.Before
	var executionPlan OneBuildExecutionPlan
	if before != "" {
		executionPlan.Init = &CommandContext{"before", before, bashCommand(sh.NewSession(), before)}
	}

	for _, name := range commands {
		executionCommand := getCommandFromName(onebuildConfig, name)
		if executionCommand == "" {
			utils.PrintlnErr("Error building exectuion plan. Command \"" + name + "\" not found.")
			config.PrintConfiguration(onebuildConfig)
			utils.Exit(1)
		}
		executionPlan.Commands = append(executionPlan.Commands, &CommandContext{name, executionCommand, bashCommand(sh.NewSession(), executionCommand)})
	}

	after := onebuildConfig.After
	if after != "" {
		executionPlan.Destroy = &CommandContext{"after", after, bashCommand(sh.NewSession(), after)}
	}

	return executionPlan
}

// OneBuildExecutionPlan holds all information for the execution strategy
type OneBuildExecutionPlan struct {
	Init     *CommandContext
	Commands []*CommandContext
	Destroy  *CommandContext
}

// CommandContext holds all meta-data and required information for execution of a command
type CommandContext struct {
	name           string
	command        string
	commandSession *sh.Session
}

func (executionPlan *OneBuildExecutionPlan) hasInit() bool {
	if executionPlan.Init != nil {
		return true
	}
	return false
}

func (executionPlan *OneBuildExecutionPlan) hasDestroy() bool {
	if executionPlan.Destroy != nil {
		return true
	}
	return false
}

func (executionPlan *OneBuildExecutionPlan) hasCommands() bool {
	if len(executionPlan.Commands) > 0 {
		return true
	}
	return false
}

func (executionPlan *OneBuildExecutionPlan) print() {
	fmt.Println()
	boldGreeen.Println("Execution plan")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.TabIndent)

	maxPhaseName := "Phase"
	maxCommand := "Command"

	if executionPlan.hasInit() {
		maxPhaseName = executionPlan.Init.name
		maxCommand = executionPlan.Init.command
	}

	if executionPlan.hasCommands() {
		for _, command := range executionPlan.Commands {
			if len(command.name) > len(maxPhaseName) {
				maxPhaseName = command.name
			}
			if len(command.command) > len(maxCommand) {
				maxCommand = command.command
			}
		}
	}

	if executionPlan.hasDestroy() {
		command := executionPlan.Destroy
		if len(command.name) > len(maxPhaseName) {
			maxPhaseName = command.name
		}
		if len(command.command) > len(maxCommand) {
			maxCommand = command.command
		}
	}

	phaseDashes := utils.DashesMatchingTextLength(maxPhaseName)
	commandDashes := utils.DashesMatchingTextLength(maxCommand)

	fmt.Fprintf(w, "%s\t%s\n", phaseDashes, commandDashes)
	fmt.Fprintln(w, "Phase\tCommand")
	fmt.Fprintf(w, "%s\t%s\n", phaseDashes, commandDashes)

	if executionPlan.hasInit() {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s", executionPlan.Init.name, executionPlan.Init.command))
	}

	if executionPlan.hasCommands() {
		for _, command := range executionPlan.Commands {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s", command.name, command.command))
		}
	}

	if executionPlan.hasDestroy() {
		fmt.Fprintln(w, fmt.Sprintf("%s\t%s", executionPlan.Destroy.name, executionPlan.Destroy.command))
	}

	w.Flush()
	fmt.Println()
	fmt.Println()
}
