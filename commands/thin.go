package commands

import (
	"fmt"
	"os"
	"strconv"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var thinCmd = &cobra.Command{
	Use:   "thin",
	Short: "Thin point cloud",
	Long:  "",
	Run:   RunThin,
}

var step, offset, limit int

func init() {
	thinCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	thinCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	thinCmd.Flags().IntVarP(&step, "step", "s", 1, "How many points to skip")
	thinCmd.Flags().IntVarP(&offset, "offset", "", 0, "Start with what point")
	thinCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Max number of points (0 = all)")
}

func RunThin(cmd *cobra.Command, args []string) {
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

	stepOpt := "--filters.decimation.step=" + strconv.Itoa(step)
	offsetOpt := "--filters.decimation.offset=" + strconv.Itoa(offset)
	limitOpt := "--filters.decimation.limit=" + strconv.Itoa(limit)

	utils.RunPdalOmni(input, output,
		"-r", readerType, "-w", writerType, "--filter", "filters.decimation",
		stepOpt, offsetOpt, limitOpt, "-v10", "--debug")

	if view {
		utils.OpenData(output)
	}
}
