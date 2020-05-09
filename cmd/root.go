package cmd

import (
	"fmt"

	"github.com/gopinath-langote/1build/cmd/config"
	parse "github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/exec"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FileFlag is a persistent (global) file flag for Root *cobra.Command
var FileFlag string

var rootCmd = &cobra.Command{
	Use:     "1build",
	Version: "1.4.0",
	Short:   "Frictionless way of managing project-specific commands",
	Args:    cobra.MinimumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := parse.LoadOneBuildConfiguration(FileFlag)
		if err != nil {
			fmt.Println(err)
			utils.ExitError()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			listCmd.Run(cmd, args)
		} else {
			fmt.Print(args, cmd)
			exec.ExecutePlan(FileFlag, args...)
		}
	},
}

// Execute entry-point for cobra app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		utils.ExitError()
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Use: "no-help", Hidden: true})
	rootCmd.PersistentFlags().StringVarP(&FileFlag, "file", "f", config.OneBuildConfigFileName, "The File Path for 1build configuration file.")
	rootCmd.PersistentFlags().BoolP("quiet", "q", false,
		"Hide output log of command & only show SUCCESS/FAILURE result")
	_ = viper.BindPFlags(rootCmd.PersistentFlags())
}
