package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var radiusCmd = &cobra.Command{
	Use: "radius",
	Run: RunRadius,
}

var minNeighbors int
var radius float64

func init() {
	radiusCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	radiusCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	radiusCmd.Flags().IntVarP(&minNeighbors, "k-neighbors", "k", 1, "Minimum number of neighbors")
	radiusCmd.Flags().Float64VarP(&radius, "radius", "r", 200.0, "Search radius")
}

func RunRadius(cmd *cobra.Command, args []string) {
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

	json := `
		{
		  "pipeline": {
		    "filters": [{
		      "name": "RadiusOutlierRemoval",
		      "setMinNeighborsRadius": ` + strconv.Itoa(minNeighbors) + `,
		      "setRadiusSearch": ` + strconv.FormatFloat(radius, 'f', -1, 64) + `
		    }]
		  }
		}
	`

	vg := `--filters.pclblock.json=` + json

	utils.RunPdalOmni(input, output,
		"-r", readerType, "-w", writerType, "--filter", "filters.pclblock",
		vg, "-v10", "--debug")

	if view {
		utils.OpenData(output)
	}
}
