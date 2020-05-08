package cmd

import (
	"fmt"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/exec"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [command..]",
	Short: "Run command(s) from the current project configuration",
	Long:  "Run command(s) from the current project configuration",
	Args:  cobra.MinimumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
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

func init() {
	rootCmd.AddCommand(runCmd)
}
