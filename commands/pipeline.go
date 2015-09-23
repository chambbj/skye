package commands

import (
	"fmt"
	"os"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Pipeline",
	Long:  "",
	Run:   RunPipeline,
}

func init() {
	pipelineCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
}

func RunPipeline(cmd *cobra.Command, args []string) {
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

	utils.RunPdalPipeline(input, "-v10", "--debug")
}
