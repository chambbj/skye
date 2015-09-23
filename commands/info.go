package commands

import (
	"fmt"
	"os"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Report point cloud info",
	Long:  "",
	Run:   RunInfo,
}

func init() {
	infoCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
}

func RunInfo(cmd *cobra.Command, args []string) {
	if input == "" {
		fmt.Println("input filename must be provided")
		cmd.Usage()
		return
	}

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", input)
		cmd.Usage()
		return
	}

	utils.RunPdal("info", input)
}
