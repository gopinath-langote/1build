package unset

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

// Cmd cobra command for unsetting/removing commands or project-level hooks from the configuration
var Cmd = &cobra.Command{
	Use:   "unset <command> [<command> ...] [--beforeAll] [--afterAll]",
	Short: "Remove one or more commands or project-level hooks from the current project configuration",
	Long: `Remove one or more commands or project-level hooks from the current project configuration.

For example:

  1build unset test lint
  1build unset --beforeAll
  1build unset --afterAll
  1build unset build --beforeAll --afterAll

This will remove the specified commands and/or project-level hooks from the configuration file.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		removeBeforeAll, _ := cmd.Flags().GetBool("beforeAll")
		removeAfterAll, _ := cmd.Flags().GetBool("afterAll")

		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			return
		}

		var notFound []string
		var removed []string

		// Remove project-level hooks if flags are set
		if removeBeforeAll {
			if configuration.BeforeAll != "" {
				configuration.BeforeAll = ""
				removed = append(removed, "beforeAll")
			}
		}
		if removeAfterAll {
			if configuration.AfterAll != "" {
				configuration.AfterAll = ""
				removed = append(removed, "afterAll")
			}
		}

		// Remove commands by name
		for _, commandName := range args {
			matched, _ := regexp.MatchString(`^[a-zA-Z0-9\-_]+$`, commandName)
			if !matched {
				fmt.Printf("1build unset: '%s' is not a valid command name. See '1build unset --help'.\n", commandName)
				utils.ExitError()
			}
			index := IndexOfCommandIfPresent(configuration, commandName)
			if index == -1 {
				notFound = append(notFound, commandName)
			} else {
				configuration.Commands = removeCommandByIndex(configuration, index)
				removed = append(removed, commandName)
			}
		}

		if len(removed) > 0 {
			err = config.WriteConfigFile(configuration)
			if err != nil {
				fmt.Println("Failed to update configuration file:", err)
				return
			}
			fmt.Printf("Removed: %s from configuration.\n", strings.Join(removed, ", "))
		}

		if len(notFound) > 0 {
			fmt.Printf("Following command(s) not found: %s\n", strings.Join(notFound, ", "))
		}
	},
}

func init() {
	Cmd.Flags().Bool("beforeAll", false, "Remove project-level beforeAll hook")
	Cmd.Flags().Bool("afterAll", false, "Remove project-level afterAll hook")
}

// IndexOfCommandIfPresent returns index in configuration for command if exists
func IndexOfCommandIfPresent(configuration config.OneBuildConfiguration, commandName string) int {
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

// removeCommandByIndex removes the command at the given index
func removeCommandByIndex(configuration config.OneBuildConfiguration, index int) []map[string]config.CommandDefinition {
	commands := configuration.Commands
	return append(commands[:index], commands[index+1:]...)
}
