package cmd

import (
	"fmt"

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
		if config.IsConfigFilePresent(FileFlag) {
			fmt.Println("'" + FileFlag + "' configuration file already exists.")
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

		err := config.WriteConfigFile(oneBuildConfiguration, FileFlag)
		if err != nil {
			fmt.Println("Failed to create file '" + FileFlag + "'")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("name", "n", "", "Project name")
	_ = initCmd.MarkFlagRequired("name")
}
