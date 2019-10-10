package cmd

import (
	"fmt"
	"regexp"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set new or update the existing command in the current project configuration",
	Long: `Set new or update the existing command in the current project configuration

- command name is a one word: without spaces, dashes and underscores are allowed
- command should be double-quoted if it the command contain spaces

For example:

  1build set test "npm run test"
  1build set npm-test "npm run test"
  1build set npm_test "npm run test"

This will update the current project configuration file.`,
	Args: cobra.ExactArgs(2),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			utils.ExitError()
		}

		commandName := args[0]
		matched, _ := regexp.MatchString(`^[a-zA-Z0-9\-_]+$`, commandName)

		if !matched {
			fmt.Println("1build set: '" + commandName + "' is not a valid command name. See '1build set --help'.")
			utils.ExitError()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		commandName := args[0]
		commandValue := args[1]

		configuration, err := buildAndSetCommand(commandName, commandValue)

		if err != nil {
			fmt.Println(err)
			return
		}

		config.WriteConfigFile(configuration)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}

func buildAndSetCommand(name string, value string) (config.OneBuildConfiguration, error) {
	configuration, err := config.LoadOneBuildConfiguration()
	if err != nil {
		return config.OneBuildConfiguration{}, err
	}

	switch name {
	case "before":
		return setBefore(configuration, value), nil
	case "after":
		return setAfter(configuration, value), nil
	default:
		return setCommand(configuration, name, value), nil
	}
}

func setBefore(configuration config.OneBuildConfiguration, value string) config.OneBuildConfiguration {
	configuration.Before = value
	return configuration
}

func setAfter(configuration config.OneBuildConfiguration, value string) config.OneBuildConfiguration {
	configuration.After = value
	return configuration
}

func setCommand(configuration config.OneBuildConfiguration, name string, value string) config.OneBuildConfiguration {
	command := map[string]string{}
	command[name] = value

	index := indexOfCommandIfPresent(configuration, name)
	if index == -1 {
		strings := append(configuration.Commands, command)
		configuration.Commands = strings
	} else {
		configuration.Commands[index] = command
	}
	return configuration
}

func indexOfCommandIfPresent(configuration config.OneBuildConfiguration, commandName string) int {
	return utils.SliceIndex(len(configuration.Commands), func(i int) bool {
		i2 := configuration.Commands[i]
		for k := range i2 {
			if k == commandName {
				return true
			}
		}
		return false
	})
}
