package commands

import (
	"fmt"
	"os"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort point cloud",
	Long:  "",
	Run:   RunSort,
}

func init() {
	sortCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	sortCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
}

func RunSort(cmd *cobra.Command, args []string) {
	if input == "" || output == "" {
		fmt.Println("input and output filenames must be provided")
		cmd.Usage()
		return
	}

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", input)
		cmd.Usage()
		return
	}

	if _, err := os.Stat(output); err == nil {
		fmt.Printf("%s exists; overwriting...\n", output)
	}

	readerType, _ := utils.InferReaderFromExt(input)
	writerType, _ := utils.InferWriterFromExt(output)

	utils.RunPdalTranslate(input, output,
		"-r", readerType, "-w", writerType, "--filter", "filters.mortonorder",
		"-v10", "--debug")

	if view {
		utils.OpenData(output)
	}
}
