package unset

import (
	"fmt"
	"github.com/gopinath-langote/1build/cmd/set"
	"regexp"
	"strings"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"

	"github.com/spf13/cobra"
)

// Cmd cobra command for unsetting one build configuration command
var Cmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove one or more existing command(s) in the current project configuration",
	Long: `Remove one or more existing command(s) in the current project configuration

- command name is a one word: without spaces, dashes and underscores are allowed

For example:

  1build unset test
  1build unset build lint other-command

This will update the current project configuration file.`,
	Args: cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			utils.ExitError()
		}

		validNameRegex := regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)
		for _, commandName := range args {
			matched := validNameRegex.MatchString(commandName)

			if !matched {
				fmt.Println("1build unset: '" + commandName + "' is not a valid command name. See '1build unset --help'.")
				utils.ExitError()
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			return
		}

		var commandsNotFound []string
		var configIsChanged bool

		for _, commandName := range args {
			index := findIndex(configuration, commandName)

			if index == -1 {
				commandsNotFound = append(commandsNotFound, commandName)
			} else {
				configuration = removeCommand(configuration, commandName, index)
				configIsChanged = true
			}
		}

		if len(commandsNotFound) != 0 {
			errorMsg := "\nFollowing command(s) not found: " + strings.Join(commandsNotFound, ", ")
			utils.CPrintln(errorMsg, utils.Style{Color: utils.RED, Bold: true})
		}

		if configIsChanged {
			_ = config.WriteConfigFile(configuration)
		}
	},
}

func removeCommandByIndex(configuration config.OneBuildConfiguration, index int) (ret []map[string]string) {
	for i, command := range configuration.Commands {
		if i != index {
			ret = append(ret, command)
		}
	}
	return
}

func findIndex(configuration config.OneBuildConfiguration, name string) int {
	switch name {
	case config.BeforeCommand, config.AfterCommand:
		return callbackExistence(configuration, name)
	default:
		return set.IndexOfCommandIfPresent(configuration, name)
	}
}

func callbackExistence(configuration config.OneBuildConfiguration, name string) int {
	switch {
	case name == config.BeforeCommand && configuration.Before == "":
		return -1
	case name == config.AfterCommand && configuration.After == "":
		return -1
	default:
		return -2
	}
}

func removeCommand(configuration config.OneBuildConfiguration, name string, index int) config.OneBuildConfiguration {
	switch name {
	case config.BeforeCommand:
		configuration.Before = ""
	case config.AfterCommand:
		configuration.After = ""
	default:
		configuration.Commands = removeCommandByIndex(configuration, index)
	}
	return configuration
}

