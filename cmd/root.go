package cmd

import (
	parse "github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/exec"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "1build",
	Version: "1.1.2",
	Short:   "Frictionless way of managing project-specific commands",
	Args:    cobra.MinimumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := parse.LoadOneBuildConfiguration()
		if err != nil {
			utils.PrintErr(err)
			utils.ExitError()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			listCmd.Run(cmd, args)
		} else {
			exec.ExecutePlan(args...)
		}
	},
}

// Execute entrypoint for cobra app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PrintErr(err)
		utils.ExitError()
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Use: "no-help", Hidden: true})
}
