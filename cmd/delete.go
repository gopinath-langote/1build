package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gopinath-langote/1build/cmd/config"
	"github.com/gopinath-langote/1build/cmd/utils"
	"github.com/spf13/cobra"
)

var shouldDelete bool

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes project configuration",
	Long: `Deletes project configuration

- To forcibly delete file without asking for consent use or --force  

For example:

  1build delete
  1build delete --force`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if !config.IsConfigFilePresent() {
			utils.CPrintln("No configuration file found!", utils.Style{Color: utils.RED})
			//fmt.Println(utils.Colored("No configuration file found!", utils.RED))
			utils.ExitError()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !shouldDelete {
			fmt.Printf("Delete 1build configuration file? (y/N) ")
			reader := bufio.NewReader(os.Stdin)
			prompt, err := reader.ReadString('\n')
			if err == nil && strings.ToLower(strings.TrimSpace(prompt)) == "y" {
				shouldDelete = true
			}
		}
		if shouldDelete {
			if err := config.DeleteConfigFile(); err != nil {
				utils.CPrintln("Error deleting configuration file.", utils.Style{Color: utils.RED})
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVar(&shouldDelete, "force", false, "Forcibly delete configuration file")
}
