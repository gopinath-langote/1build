package exec

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/codeskyblue/go-sh"
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/models"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/viper"
)

// ExecutePlan executes the Execution plan
func ExecutePlan(commands ...string) {
	executeStart := time.Now()

	configuration, err := config.LoadOneBuildConfiguration()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		utils.ExitError()
	}

	// Handle SIGINT / SIGTERM: print FAILURE banner before exiting.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Fprintln(os.Stderr, "\nInterrupted.")
		printResultsBanner(false, executeStart)
		os.Exit(1)
	}()

	// Print and execute global before-all
	if configuration.BeforeAll != "" {
		fmt.Printf("before-all: %s\n", configuration.BeforeAll)
		executeAndStopIfFailed(&models.CommandContext{
			Name:           "before-all",
			Command:        configuration.BeforeAll,
			CommandSession: bashCommand(sh.NewSession(), configuration.BeforeAll),
		}, executeStart)
	}

	// For each command argument, find and execute with hooks
	for _, name := range commands {
		var def config.CommandDefinition
		found := false
		for _, cmdMap := range configuration.Commands {
			if d, ok := cmdMap[name]; ok {
				def = d
				found = true
				break
			}
		}
		if !found {
			utils.CPrintlnErr("\nError: Command \""+name+"\" not found.", utils.Style{Color: utils.RED, Bold: true})
			configuration.Print()
			utils.ExitWithCode("127")
		}

		// Command-level before
		if def.Before != "" {
			fmt.Printf("  before: %s\n", def.Before)
			executeAndStopIfFailed(&models.CommandContext{
				Name:           name + ":before",
				Command:        def.Before,
				CommandSession: bashCommand(sh.NewSession(), def.Before),
			}, executeStart)
		}

		// Main command
		fmt.Printf("Executing command: %s\n", name)
		if def.Command != "" {
			fmt.Printf("  command: %s\n", def.Command)
		}
		executeAndStopIfFailed(&models.CommandContext{
			Name:           name,
			Command:        def.Command,
			CommandSession: bashCommand(sh.NewSession(), def.Command),
		}, executeStart)

		// Command-level after
		if def.After != "" {
			fmt.Printf("  after: %s\n", def.After)
			executeAndStopIfFailed(&models.CommandContext{
				Name:           name + ":after",
				Command:        def.After,
				CommandSession: bashCommand(sh.NewSession(), def.After),
			}, executeStart)
		}
	}

	// Print and execute global after-all
	if configuration.AfterAll != "" {
		fmt.Printf("after-all: %s\n", configuration.AfterAll)
		executeAndStopIfFailed(&models.CommandContext{
			Name:           "after-all",
			Command:        configuration.AfterAll,
			CommandSession: bashCommand(sh.NewSession(), configuration.AfterAll),
		}, executeStart)
	}

	printResultsBanner(true, executeStart)
}

func executeAndStopIfFailed(command *models.CommandContext, executeStart time.Time) {
	command.PrintPhaseBanner()
	session := command.CommandSession
	session.SetStdin(os.Stdin)

	if !viper.GetBool("quiet") {
		err := session.Run()
		if err != nil {
			exitCode := strings.TrimPrefix(err.Error(), "exit status ")
			text := "\nExecution failed in phase '" + command.Name + "' – exit code: " + exitCode
			utils.CPrintlnErr(text, utils.Style{Color: utils.RED})
			printResultsBanner(false, executeStart)
			utils.ExitWithCode(exitCode)
		}
	} else {
		_, err := session.CombinedOutput()
		if err != nil {
			exitCode := strings.TrimPrefix(err.Error(), "exit status ")
			text := "\nExecution failed in phase '" + command.Name + "' – exit code: " + exitCode
			utils.CPrintlnErr(text, utils.Style{Color: utils.RED})
			printResultsBanner(false, executeStart)
			utils.ExitWithCode(exitCode)
		}
	}
}

func bashCommand(s *sh.Session, command string) *sh.Session {
	configFileAbsoluteDir, _ := config.GetAbsoluteDirPathOfConfigFile()
	s.SetDir(configFileAbsoluteDir)
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
	fmt.Println()
	fmt.Println(utils.Dash())
	if isSuccess {
		utils.CPrint("SUCCESS", utils.Style{Color: utils.CYAN, Bold: true})
	} else {
		utils.CPrint("FAILURE", utils.Style{Color: utils.RED, Bold: true})
	}
	fmt.Printf(" - Total Time: %s\n", timeStr)
	fmt.Println(utils.Dash())
}
