package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var groundCmd = &cobra.Command{
	Use:   "ground",
	Short: "Segment ground returns",
	Long:  "",
	Run:   RunGround,
}

var winSize int
var slope, maxDist, initDist, cellSize, base float64
var isExp, isNeg bool

func init() {
	groundCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	groundCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	groundCmd.Flags().IntVarP(&winSize, "winSize", "", 33, "Max window size")
	groundCmd.Flags().Float64VarP(&slope, "slope", "", 1.0, "Slope")
	groundCmd.Flags().Float64VarP(&maxDist, "maxDist", "", 2.5, "Max distance")
	groundCmd.Flags().Float64VarP(&initDist, "initDist", "", 0.15, "Initial distance")
	groundCmd.Flags().Float64VarP(&cellSize, "cellSize", "", 1.0, "Cell size")
	groundCmd.Flags().Float64VarP(&base, "base", "", 2.0, "Base")
	groundCmd.Flags().BoolVarP(&isExp, "isExp", "", true, "Exponential?")
	groundCmd.Flags().BoolVarP(&isNeg, "isNeg", "", false, "Negative?")
}

func RunGround(cmd *cobra.Command, args []string) {
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
		      "name": "ProgressiveMorphologicalFilter",
		      "setMaxWindowSize": ` + strconv.Itoa(winSize) + `,
		      "setSlope": ` + strconv.FormatFloat(slope, 'f', -1, 64) + `,
		      "setMaxDistance": ` + strconv.FormatFloat(maxDist, 'f', -1, 64) + `,
		      "setInitialDistance": ` + strconv.FormatFloat(initDist, 'f', -1, 64) + `,
		      "setCellSize": ` + strconv.FormatFloat(cellSize, 'f', -1, 64) + `,
		      "setBase": ` + strconv.FormatFloat(base, 'f', -1, 64) + `,
		      "setExponential": ` + strconv.FormatBool(isExp) + `,
		      "setNegative": ` + strconv.FormatBool(isNeg) + `
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
