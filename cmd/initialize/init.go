package initialize

import (
	"fmt"
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

// Cmd cobra command for initializing one build configuration
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Create default project configuration",
	Long: `Create default project configuration

- Name of the project needs be passed as parameter - 'name' 

For example:

  1build initialize --name project
  1build initialize --name "My favorite project"`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if config.IsConfigFilePresent() {
			fmt.Println("'" + config.OneBuildConfigFileName + "' configuration file already exists.")
			utils.ExitError()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("name")

		defaultCommand := map[string]string{}
		defaultCommand["build"] = "echo 'Running build'"

		oneBuildConfiguration := config.OneBuildConfiguration{
			Project:  projectName,
			Commands: []map[string]string{defaultCommand},
		}

		err := config.WriteConfigFile(oneBuildConfiguration)
		if err != nil {
			fmt.Println("Failed to create file '" + config.OneBuildConfigFileName + "'")
		}
	},
}

func init() {
	Cmd.Flags().StringP("name", "n", "", "Project name")
	_ = Cmd.MarkFlagRequired("name")
}
