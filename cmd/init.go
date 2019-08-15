package cmd

import (
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create default project configuration",
	Long: `Create default project configuration

- Name of the project needs be passed as parameter - 'name' 

For example:

  1build init --name project
  1build init --name "My favorite project"`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if config.IsConfigFilePresent() {
			utils.Println("'" + config.OneBuildConfigFileName + "' configuration file already exists.")
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
			utils.Println("Failed to create file '" + config.OneBuildConfigFileName + "'")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("name", "n", "", "Project name")
	_ = initCmd.MarkFlagRequired("name")
}
