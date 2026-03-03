package set

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

// Cmd cobra command for setting one build configuration command or project-level hooks
var Cmd = &cobra.Command{
	Use: "set <name> [command] [--command <command>] [--before <before>] [--after <after>] " +
		"[--before-all <cmd>] [--after-all <cmd>]",
	Short: "Set or update a command or project-level hooks in the current project configuration",
	Long: `Set or update a command or project-level hooks in the current project configuration.

You have three options:

1. Set a simple (vanilla) command (inline form):
   1build set test "npm run test"
   1build set test --command "npm run test"

2. Set a nested command with before/after hooks:
   1build set build --command "npm run build" --before "echo before" --after "echo after"

3. Set project-level hooks (--before-all/--after-all):
   1build set --before-all "echo before all"
   1build set --after-all "echo after all"
   1build set --before-all "echo before all" --after-all "echo after all"

- <name> is a single word: no spaces, dashes and underscores are allowed.
- Command can be provided as a positional argument or with --command flag.
- --before and --after are optional for command-level hooks.
- --before-all and --after-all set project-level hooks.
- Use --dry-run to preview changes without writing to disk.

This will update the current project configuration file.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		beforeAll, _ := cmd.Flags().GetString("before-all")
		afterAll, _ := cmd.Flags().GetString("after-all")
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			utils.ExitError()
		}

		var updated []string

		// Update project-level hooks if flags are set
		if beforeAll != "" {
			configuration.BeforeAll = beforeAll
			updated = append(updated, "before-all")
		}
		if afterAll != "" {
			configuration.AfterAll = afterAll
			updated = append(updated, "after-all")
		}

		// Command-level logic
		if len(args) > 0 && args[0] != "" {
			commandName := args[0]

			matched, _ := regexp.MatchString(`^[a-zA-Z0-9\-_]+$`, commandName)
			if !matched {
				fmt.Fprintln(os.Stderr, "1build set: '"+commandName+"' is not a valid command name. See '1build set --help'.")
				utils.ExitUsage()
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
				fmt.Fprintln(os.Stderr, "Error: command is required as a positional argument or with --command flag. See '1build set --help'.")
				utils.ExitUsage()
			}

			def := config.CommandDefinition{
				Before:  before,
				Command: command,
				After:   after,
			}

			cmdMap := map[string]config.CommandDefinition{
				commandName: def,
			}

			index := configuration.IndexOfCommand(commandName)
			if index == -1 {
				configuration.Commands = append(configuration.Commands, cmdMap)
			} else {
				configuration.Commands[index] = cmdMap
			}
			updated = append(updated, fmt.Sprintf("command '%s'", commandName))
		}

		if len(updated) == 0 {
			fmt.Println("No changes made to configuration.")
			return
		}

		if dryRun {
			fmt.Printf("[dry-run] Would update: %s in configuration.\n", strings.Join(updated, ", "))
			return
		}

		err = config.WriteConfigFile(configuration)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to update configuration file:", err)
			utils.ExitError()
		}

		fmt.Printf("Updated: %s in configuration.\n", strings.Join(updated, ", "))
	},
}
