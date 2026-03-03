package list

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

// Cmd cobra command for listing all commands in the configuration
var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List all commands in the current project configuration",
	Long: `List all commands in the current project configuration, including any before/after hooks.

Examples:

  1build list
  1build list --output json`,
	Run: func(cmd *cobra.Command, args []string) {
		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			utils.ExitError()
		}

		outputFormat, _ := cmd.Flags().GetString("output")

		if outputFormat == "json" {
			printJSON(configuration)
			return
		}

		// Default human-readable output
		configuration.Print()
	},
}

func init() {
	Cmd.Flags().StringP("output", "o", "", "Output format: json")
}

// printJSON serialises the configuration to JSON and writes it to stdout.
func printJSON(configuration config.OneBuildConfiguration) {
	type commandEntry struct {
		Name    string `json:"name"`
		Before  string `json:"before,omitempty"`
		Command string `json:"command,omitempty"`
		After   string `json:"after,omitempty"`
	}
	type output struct {
		Project   string         `json:"project"`
		BeforeAll string         `json:"beforeAll,omitempty"`
		AfterAll  string         `json:"afterAll,omitempty"`
		Commands  []commandEntry `json:"commands"`
	}

	var cmds []commandEntry
	for _, cmdMap := range configuration.Commands {
		for name, def := range cmdMap {
			cmds = append(cmds, commandEntry{
				Name:    name,
				Before:  def.Before,
				Command: def.Command,
				After:   def.After,
			})
		}
	}

	out := output{
		Project:   configuration.Project,
		BeforeAll: configuration.BeforeAll,
		AfterAll:  configuration.AfterAll,
		Commands:  cmds,
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(out); err != nil {
		fmt.Fprintln(os.Stderr, "Error encoding JSON:", err)
		utils.ExitError()
	}
}
