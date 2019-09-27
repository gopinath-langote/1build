package cmd

import (
	"regexp"
	"strings"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"

	"github.com/spf13/cobra"
)

var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove one or more existing command(s) in the current project configuration",
	Long: `Remove one or more existing command(s) in the current project configuration

- command name is a one word: without spaces, dashes and underscores are allowed

For example:

  1build unset test build

This will update the current project configuration file.`,
	Args: cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := config.LoadOneBuildConfiguration()
		if err != nil {
			utils.PrintErr(err)
			utils.ExitError()
		}

		commandName := args[0]
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9\-_]+$`, commandName)

		if !matched {
			utils.Println("1build unset: '" + commandName + "' is not a valid command name. See '1build unset --help'.")
			utils.ExitError()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			utils.PrintErr(err)
			return
		}

		var commandsNotFound []string
		var configIsChanged bool

		for _, commandName := range args {
			index := indexOfCommandIfPresent(configuration, commandName)
			if index == -1 {
				commandsNotFound = append(commandsNotFound, commandName)
			} else {
				configuration.Commands = removeCommandByIndex(configuration, index)
				configIsChanged = true
			}
		}

		if len(commandsNotFound) != 0 {
			errorMsg := "Following command(s) not found: " + strings.Join(commandsNotFound, ", ")
			utils.PrintlnErr(errorMsg)
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

func init() {
	rootCmd.AddCommand(unsetCmd)
}
