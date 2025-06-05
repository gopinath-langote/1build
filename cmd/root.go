package cmd

import (
	"fmt"

	"github.com/gopinath-langote/1build/cmd/del"
	"github.com/gopinath-langote/1build/cmd/initialize"
	"github.com/gopinath-langote/1build/cmd/list"
	"github.com/gopinath-langote/1build/cmd/set"
	"github.com/gopinath-langote/1build/cmd/unset"

	configuration "github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/exec"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cmd cobra for root level
var Cmd = &cobra.Command{
	Use:     "1build",
	Version: "1.4.0",
	Short:   "Frictionless way of managing project-specific commands",
	Args:    cobra.MinimumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := configuration.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			utils.ExitError()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			list.Cmd.Run(cmd, args)
		} else {
			exec.ExecutePlan(args...)
		}
	},
}

// Execute entry-point for cobra app
func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		utils.ExitError()
	}
}

func init() {
	Cmd.SetHelpCommand(&cobra.Command{Use: "no-help", Hidden: true})
	Cmd.PersistentFlags().BoolP("quiet", "q", false,
		"Hide output log of command & only show SUCCESS/FAILURE result")
	Cmd.PersistentFlags().
		StringP("file", "f", configuration.OneBuildConfigFileName, "The file path for 1build configuration file.")
	_ = viper.BindPFlags(Cmd.PersistentFlags())

	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(del.Cmd)
	Cmd.AddCommand(initialize.Cmd)
	Cmd.AddCommand(set.Cmd)
	Cmd.AddCommand(unset.Cmd)

	// set command flags
	set.Cmd.Flags().String("before", "", "Command to execute before the main command")
	set.Cmd.Flags().String("command", "", "Main command to execute (can also be provided as a positional argument)")
	set.Cmd.Flags().String("after", "", "Command to execute after the main command")
	set.Cmd.Flags().String("beforeAll", "", "Project-level command to execute before all commands")
	set.Cmd.Flags().String("afterAll", "", "Project-level command to execute after all commands")

	// unset command flags
	unset.Cmd.Flags().Bool("beforeAll", false, "Remove project-level beforeAll hook")
	unset.Cmd.Flags().Bool("afterAll", false, "Remove project-level afterAll hook")

	// del command flags
	del.Cmd.Flags().Bool("force", false, "Force delete command")

	// init command flags
	initialize.Cmd.Flags().String("name", "", "Project name")
	_ = Cmd.MarkFlagRequired("name")
}
