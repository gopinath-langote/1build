package rename

import (
	"fmt"
	"os"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

// Cmd cobra command for renaming a command in the project configuration
var Cmd = &cobra.Command{
	Use:   "rename <old-name> <new-name>",
	Short: "Rename a command in the project configuration",
	Long: `Rename an existing command in the project configuration.

For example:

  1build rename build compile
  1build rename test verify`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		oldName := args[0]
		newName := args[1]

		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			utils.ExitError()
		}

		oldIndex := configuration.IndexOfCommand(oldName)
		if oldIndex == -1 {
			fmt.Fprintf(os.Stderr, "Error: command '%s' not found in configuration.\n", oldName)
			utils.ExitError()
		}

		if configuration.IndexOfCommand(newName) != -1 {
			fmt.Fprintf(os.Stderr, "Error: command '%s' already exists in configuration.\n", newName)
			utils.ExitError()
		}

		// Replace the key in the map at the found index while preserving the definition.
		oldDef := configuration.Commands[oldIndex][oldName]
		configuration.Commands[oldIndex] = map[string]config.CommandDefinition{newName: oldDef}

		if err := config.WriteConfigFile(configuration); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to update configuration file:", err)
			utils.ExitError()
		}

		fmt.Printf("Renamed command '%s' to '%s'.\n", oldName, newName)
	},
}
