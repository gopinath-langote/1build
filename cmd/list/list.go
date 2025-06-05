package list

import (
	"fmt"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/spf13/cobra"
)

// Cmd cobra command for listing all commands in the configuration
var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List all commands in the current project configuration",
	Long:  "List all commands in the current project configuration, including any before/after hooks.",
	Run: func(cmd *cobra.Command, args []string) {
		configuration, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Project: %s\n", configuration.Project)
		if configuration.BeforeAll != "" {
			fmt.Printf("beforeAll: %s\n", configuration.BeforeAll)
		}
		if configuration.AfterAll != "" {
			fmt.Printf("afterAll: %s\n", configuration.AfterAll)
		}
		fmt.Println("commands:")
		for _, command := range configuration.Commands {
			for name, def := range command {
				// Print nested structure if present
				if def.Before != "" || def.After != "" {
					fmt.Printf("  %s:\n", name)
					if def.Before != "" {
						fmt.Printf("    before: %s\n", def.Before)
					}
					if def.Command != "" {
						fmt.Printf("    command: %s\n", def.Command)
					}
					if def.After != "" {
						fmt.Printf("    after: %s\n", def.After)
					}
				} else {
					fmt.Printf("  %s: %s\n", name, def.Command)
				}
			}
		}
	},
}
