package commands

import (
	"fmt"
	"os"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var cropCmd = &cobra.Command{
	Use:   "crop",
	Short: "Crop point cloud",
	Long:  "",
	Run:   RunCrop,
}

var bounds, polygon string
var outside bool

func init() {
	cropCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	cropCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	cropCmd.Flags().StringVarP(&bounds, "bounds", "b", "", "Bounds")
	cropCmd.Flags().StringVarP(&polygon, "polygon", "p", "", "Polygon")
	cropCmd.Flags().BoolVarP(&outside, "outside", "", false, "Only take points outside the bounds or polygon")
}

func RunCrop(cmd *cobra.Command, args []string) {
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

	if (bounds == "" && polygon == "") || (bounds != "" && polygon != "") {
		fmt.Println("must provide bounds OR polygon, but not both")
	}

	var geometry string
	if bounds != "" {
		geometry = "--filters.crop.bounds=" + bounds
	} else if polygon != "" {
		geometry = "--filters.crop.polygon=" + polygon
	}

	var invert string
	if outside {
		invert = "--filters.crop.outside=true"
	} else {
		invert = "--filters.crop.outside=false"
	}

	utils.RunPdalOmni(input, output,
		"-r", readerType, "-w", writerType, "--filter", "filters.crop",
		geometry, invert, "-v10", "--debug")

	if view {
		utils.OpenData(output)
	}
}
