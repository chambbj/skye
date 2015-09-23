package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var statisticalCmd = &cobra.Command{
	Use: "statistical",
	Run: RunStatistical,
}

var meanK int
var thresh float64

func init() {
	statisticalCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	statisticalCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	statisticalCmd.Flags().IntVarP(&meanK, "k-neighbors", "k", 2, "Mean number of neighbors")
	statisticalCmd.Flags().Float64VarP(&thresh, "thresh", "t", 1.5, "Standard deviation multiplier threshold")
}

func RunStatistical(cmd *cobra.Command, args []string) {
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
		      "name": "StatisticalOutlierRemoval",
		      "setMinNeighborsRadius": ` + strconv.Itoa(meanK) + `,
		      "setRadiusSearch": ` + strconv.FormatFloat(thresh, 'f', -1, 64) + `
		    }]
		  }
		}
	`

	vg := `--filters.pclblock.json=` + json

	utils.RunPdalTranslate(input, output,
		"-r", readerType, "-w", writerType, "--filter", "filters.pclblock",
		vg, "-v10", "--debug")

	if view {
		utils.OpenData(output)
	}
}
