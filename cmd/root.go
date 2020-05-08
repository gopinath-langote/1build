package cmd

import (
	"fmt"

	parse "github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "1build",
	Version: "1.4.0",
	Short:   "Frictionless way of managing project-specific commands",
	Args:    cobra.MinimumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		_, err := parse.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			utils.ExitError()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a command to 1build")
		cmd.Help()
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
	rootCmd.PersistentFlags().BoolP("quiet", "q", false,
		"Hide output log of command & only show SUCCESS/FAILURE result")
	_ = viper.BindPFlags(rootCmd.PersistentFlags())
}
