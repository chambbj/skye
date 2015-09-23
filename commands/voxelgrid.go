package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var voxelGridCmd = &cobra.Command{
	Use:   "voxelgrid",
	Short: "Decimate point cloud",
	Long:  "",
	Run:   RunVoxelGrid,
}

var gridX, gridY, gridZ float64

func init() {
	voxelGridCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	voxelGridCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	voxelGridCmd.Flags().Float64VarP(&gridX, "gridX", "", 5.0, "Grid X")
	voxelGridCmd.Flags().Float64VarP(&gridY, "gridY", "", 5.0, "Grid Y")
	voxelGridCmd.Flags().Float64VarP(&gridZ, "gridZ", "", 5.0, "Grid Z")
}

func RunVoxelGrid(cmd *cobra.Command, args []string) {
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
		      "name": "VoxelGrid",
		      "setLeafSize":
		      {
		        "x": ` + strconv.FormatFloat(gridX, 'f', -1, 64) + `,
		        "y": ` + strconv.FormatFloat(gridY, 'f', -1, 64) + `,
		        "z": ` + strconv.FormatFloat(gridZ, 'f', -1, 64) + `
		      }
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
