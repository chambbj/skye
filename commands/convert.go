package commands

import (
	"fmt"
	"os"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert point cloud formats",
	Long:  "",
	Run:   RunConvert,
}

func init() {
	convertCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	convertCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	convertCmd.Flags().StringVarP(&reader, "reader", "r", "", "Reader type")
	convertCmd.Flags().StringVarP(&writer, "writer", "w", "", "Writer type")
}

// RunConvert runs a PDAL pipeline with colorization filter.
func RunConvert(cmd *cobra.Command, args []string) {
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

	if reader == "" {
		reader, _ = utils.InferReaderFromExt(input)
	}
	if writer == "" {
		writer, _ = utils.InferWriterFromExt(output)
	}

	utils.RunPdalTranslate(input, output,
		"-r", reader, "-w", writer, "-v10", "--debug")

	if view {
		utils.OpenData(output)
	}
}
