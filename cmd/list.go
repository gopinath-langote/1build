package cmd

import (
	"fmt"
	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all available commands from the current project configuration",
	Long:  "Show all available commands from the current project configuration",
	Run: func(cmd *cobra.Command, args []string) {
		oneBuildConfig, err := config.LoadOneBuildConfiguration()
		if err != nil {
			fmt.Println(err)
			return
		}
		oneBuildConfig.Print()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
