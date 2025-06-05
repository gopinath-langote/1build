package set

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

// Cmd cobra command for setting one build configuration command or project-level hooks
var Cmd = &cobra.Command{
	Use:   "set <name> [command] [--command <command>] [--before <before>] [--after <after>] [--beforeAll <beforeAll>] [--afterAll <afterAll>]",
	Short: "Set or update a command or project-level hooks in the current project configuration",
	Long: `Set or update a command or project-level hooks in the current project configuration.

You have three options:

1. Set a simple (vanilla) command (inline form):
   1build set test "npm run test"
   1build set test --command "npm run test"

2. Set a nested command with before/after hooks:
   1build set build --command "npm run build" --before "echo before" --after "echo after"

3. Set project-level hooks (beforeAll/afterAll):
   1build set --beforeAll "echo before all"
   1build set --afterAll "echo after all"
   1build set --beforeAll "echo before all" --afterAll "echo after all"

- <name> is a single word: no spaces, dashes and underscores are allowed.
- Command can be provided as a positional argument or with --command flag.
- --before and --after are optional for command-level hooks.
- --beforeAll and --afterAll set project-level hooks.
- If you only want a simple command, just use the positional argument or --command.

This will update the current project configuration file.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		beforeAll, _ := cmd.Flags().GetString("beforeAll")
		afterAll, _ := cmd.Flags().GetString("afterAll")

		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			return
		}

		var updated []string

		// Update project-level hooks if flags are set
		if beforeAll != "" {
			configuration.BeforeAll = beforeAll
			updated = append(updated, "beforeAll")
		}
		if afterAll != "" {
			configuration.AfterAll = afterAll
			updated = append(updated, "afterAll")
		}

		// Command-level logic
		if len(args) > 0 && args[0] != "" {
			commandName := args[0]

			matched, _ := regexp.MatchString(`^[a-zA-Z0-9\-_]+$`, commandName)
			if !matched {
				fmt.Println("1build set: '" + commandName + "' is not a valid command name. See '1build set --help'.")
				utils.ExitError()
			}

			before, _ := cmd.Flags().GetString("before")
			commandFlag, _ := cmd.Flags().GetString("command")
			after, _ := cmd.Flags().GetString("after")

			// Prefer --command flag, fallback to positional argument if present
			command := commandFlag
			if command == "" && len(args) > 1 {
				command = args[1]
			}

			if command == "" {
				fmt.Println("Error: command is required as a positional argument or with --command flag. See '1build set --help'.")
				utils.ExitError()
			}

			def := config.CommandDefinition{
				Before:  before,
				Command: command,
				After:   after,
			}

			cmdMap := map[string]config.CommandDefinition{
				commandName: def,
			}

			index := IndexOfCommandIfPresent(configuration, commandName)
			if index == -1 {
				configuration.Commands = append(configuration.Commands, cmdMap)
			} else {
				configuration.Commands[index] = cmdMap
			}
			updated = append(updated, fmt.Sprintf("command '%s'", commandName))
		}

		err = config.WriteConfigFile(configuration)
		if err != nil {
			fmt.Println("Failed to update configuration file:", err)
			return
		}

		if len(updated) > 0 {
			fmt.Printf("Updated: %s in configuration.\n", strings.Join(updated, ", "))
		} else {
			fmt.Println("No changes made to configuration.")
		}
	},
}

func init() {
	Cmd.Flags().String("before", "", "Command to execute before the main command")
	Cmd.Flags().String("command", "", "Main command to execute (can also be provided as a positional argument)")
	Cmd.Flags().String("after", "", "Command to execute after the main command")
	Cmd.Flags().String("beforeAll", "", "Project-level command to execute before all commands")
	Cmd.Flags().String("afterAll", "", "Project-level command to execute after all commands")
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
