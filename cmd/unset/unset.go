package unset

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

// Cmd cobra command for unsetting/removing commands or project-level hooks from the configuration
var Cmd = &cobra.Command{
	Use:   "unset <command> [<command> ...] [--before-all] [--after-all]",
	Short: "Remove one or more commands or project-level hooks from the current project configuration",
	Long: `Remove one or more commands or project-level hooks from the current project configuration.

For example:

  1build unset test lint
  1build unset --before-all
  1build unset --after-all
  1build unset build --before-all --after-all

Use --dry-run to preview changes without writing to disk.

This will remove the specified commands and/or project-level hooks from the configuration file.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		removeBeforeAll, _ := cmd.Flags().GetBool("before-all")
		removeAfterAll, _ := cmd.Flags().GetBool("after-all")
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			utils.ExitError()
		}

		var notFound []string
		var removed []string

		// Remove project-level hooks if flags are set
		if removeBeforeAll {
			if configuration.BeforeAll != "" {
				configuration.BeforeAll = ""
				removed = append(removed, "before-all")
			}
		}
		if removeAfterAll {
			if configuration.AfterAll != "" {
				configuration.AfterAll = ""
				removed = append(removed, "after-all")
			}
		}

		reValidName := regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)
		// Remove commands by name
		for _, commandName := range args {
			if !reValidName.MatchString(commandName) {
				fmt.Fprintf(os.Stderr, "1build unset: '%s' is not a valid command name. See '1build unset --help'.\n", commandName)
				utils.ExitUsage()
			}
			index := configuration.IndexOfCommand(commandName)
			if index == -1 {
				notFound = append(notFound, commandName)
			} else {
				configuration.Commands = removeCommandByIndex(configuration, index)
				removed = append(removed, commandName)
			}
		}

		if len(notFound) > 0 {
			fmt.Printf("Following command(s) not found: %s\n", strings.Join(notFound, ", "))
		}

		if len(removed) > 0 {
			if dryRun {
				fmt.Printf("[dry-run] Would remove: %s from configuration.\n", strings.Join(removed, ", "))
				return
			}
			err = config.WriteConfigFile(configuration)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to update configuration file:", err)
				utils.ExitError()
			}
			fmt.Printf("Removed: %s from configuration.\n", strings.Join(removed, ", "))
		}
	},
}

// removeCommandByIndex removes the command at the given index
func removeCommandByIndex(configuration config.OneBuildConfiguration, index int) []map[string]config.CommandDefinition {
	commands := configuration.Commands
	return append(commands[:index], commands[index+1:]...)
}
