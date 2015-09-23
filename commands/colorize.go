package commands

import (
	"fmt"
	"os"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var raster string

var colorizeCmd = &cobra.Command{
	Use:   "colorize",
	Short: "Colorize point cloud",
	Long:  "",
	Run:   RunColorize,
}

func init() {
	colorizeCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	colorizeCmd.Flags().StringVarP(&raster, "raster", "r", "", "Raster filename")
	colorizeCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
}

// RunColorize runs a PDAL pipeline with colorization filter.
func RunColorize(cmd *cobra.Command, args []string) {
	if input == "" || output == "" || raster == "" {
		fmt.Println("input, raster, and output filenames must be provided")
		cmd.Usage()
		return
	}

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", input)
		cmd.Usage()
		return
	}

	if _, err := os.Stat(raster); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", raster)
		cmd.Usage()
		return
	}

	if _, err := os.Stat(output); err == nil {
		fmt.Printf("%s exists; overwriting...\n", output)
	}

	readerType, _ := utils.InferReaderFromExt(input)
	writerType, _ := utils.InferWriterFromExt(output)

	rasterOpt := "--filters.colorization.raster=" + raster

	utils.RunPdalOmni(input, output,
		"-r", readerType, "-w", writerType, "--filter", "filters.colorization",
		rasterOpt, "-v10", "--debug")

	if view {
		utils.OpenData(output)
	}
}
