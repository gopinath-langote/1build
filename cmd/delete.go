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

- To forcibly delete file without asking for consent use -f or --force  

For example:

  1build delete
  1build delet --force`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if !config.IsConfigFilePresent() {
			fmt.Println("No configuration file found!")
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
				fmt.Println("Error deleting configuration file.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&shouldDelete, "force", "f", false, "Forcibly delete configuration file")
}
