package cmd

import (
	"fmt"
	"os"

	"github.com/gopinath-langote/1build/cmd/completion"
	"github.com/gopinath-langote/1build/cmd/del"
	"github.com/gopinath-langote/1build/cmd/initialize"
	"github.com/gopinath-langote/1build/cmd/list"
	"github.com/gopinath-langote/1build/cmd/rename"
	"github.com/gopinath-langote/1build/cmd/set"
	"github.com/gopinath-langote/1build/cmd/unset"

	configuration "github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/exec"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Version is set at build time via -ldflags. Defaults to "dev" for local builds.
var Version = "dev"

// Cmd cobra for root level
var Cmd = &cobra.Command{
	Use:     "1build",
	Version: Version,
	Short:   "Frictionless way of managing project-specific commands",
	Args:    cobra.MinimumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := configuration.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
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
		fmt.Fprintln(os.Stderr, err)
		utils.ExitUsage()
	}
}

func init() {
	Cmd.PersistentFlags().BoolP("quiet", "q", false,
		"Hide output log of command & only show SUCCESS/FAILURE result")
	Cmd.PersistentFlags().
		StringP("file", "f", configuration.OneBuildConfigFileName, "The file path for 1build configuration file.")
	_ = viper.BindPFlags(Cmd.PersistentFlags())

	// Register -v as an alias for --version
	Cmd.Flags().BoolP("version", "v", false, "version for 1build")

	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(del.Cmd)
	Cmd.AddCommand(initialize.Cmd)
	Cmd.AddCommand(set.Cmd)
	Cmd.AddCommand(unset.Cmd)
	Cmd.AddCommand(rename.Cmd)
	Cmd.AddCommand(completion.Cmd)

	// set command flags
	set.Cmd.Flags().String("before", "", "Command to execute before the main command")
	set.Cmd.Flags().String("command", "", "Main command to execute (can also be provided as a positional argument)")
	set.Cmd.Flags().String("after", "", "Command to execute after the main command")
	set.Cmd.Flags().String("before-all", "", "Project-level command to execute before all commands")
	set.Cmd.Flags().String("after-all", "", "Project-level command to execute after all commands")
	set.Cmd.Flags().Bool("dry-run", false, "Preview changes without writing to disk")

	// unset command flags
	unset.Cmd.Flags().Bool("before-all", false, "Remove project-level beforeAll hook")
	unset.Cmd.Flags().Bool("after-all", false, "Remove project-level afterAll hook")
	unset.Cmd.Flags().Bool("dry-run", false, "Preview changes without writing to disk")

	// del command flags
	del.Cmd.Flags().Bool("force", false, "Force delete command")
	del.Cmd.Flags().Bool("dry-run", false, "Preview what would be deleted without deleting")

	// rename command flags
	rename.Cmd.Flags().Bool("dry-run", false, "Preview changes without writing to disk")

	// init command flags
	initialize.Cmd.Flags().String("name", "", "Project name to use in the generated 1build.yaml")
}
