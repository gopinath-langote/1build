package initialize

import (
	"fmt"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/spf13/cobra"
)

// Cmd cobra command for initializing one build configuration
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize 1build configuration file in the current directory",
	Long: `Initialize 1build configuration file in the current directory.

This will create a sample 1build.yaml file if it does not exist.

Use --name to set the project name in the generated configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.IsConfigFilePresent() {
			fmt.Printf("'%s' already exists in the current directory.\n", config.OneBuildConfigFileName)
			return
		}

		// Use --name flag if provided, fall back to default
		projectName, _ := cmd.Flags().GetString("name")
		if projectName == "" {
			projectName = "Sample Project"
		}

		configuration := config.OneBuildConfiguration{
			Project: projectName,
			Commands: []map[string]config.CommandDefinition{
				{"setup": {Command: "go get -u golang.org/x/lint/golint"}},
				{"test": {Command: "go test ./..."}},
				{"lint": {Command: "golint ./..."}},
				{"build": {
					Before:  "echo \"before build\"",
					Command: "go build",
					After:   "echo \"after build\"",
				}},
			},
		}

		err := config.WriteConfigFile(configuration)
		if err != nil {
			fmt.Println("Failed to create configuration file:", err)
			return
		}

		fmt.Printf("Created '%s' in the current directory.\n", config.OneBuildConfigFileName)
	},
}
