package cmd

import (
	"regexp"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"

	"github.com/spf13/cobra"
)

var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove the existing command in the current project configuration",
	Long: `Remove the existing command in the current project configuration

- command name is a one word: without spaces, dashes and underscores are allowed

For example:

  1build unset test

This will update the current project configuration file.`,
	Args: cobra.ExactArgs(1),
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
		commandName := args[0]

		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			utils.PrintErr(err)
			return
		}

		index := indexOfCommandIfPresent(configuration, commandName)
		if index == -1 {
			utils.Println("Command '" + commandName + "' not found")
			return
		}

		configuration.Commands = removeCommandByIndex(configuration, index)
		_ = config.WriteConfigFile(configuration)
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
