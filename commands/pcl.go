package commands

import (
	"fmt"
	"os"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var pclCmd = &cobra.Command{
	Use:   "pcl",
	Short: "Invoke PCL block",
	Long:  "",
	Run:   RunPCL,
}

var pcl string

func init() {
	pclCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
	pclCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
	pclCmd.Flags().StringVarP(&pcl, "pcl", "p", "", "PCL filename")
}

func RunPCL(cmd *cobra.Command, args []string) {
	if input == "" || output == "" || pcl == "" {
		fmt.Println("input, output, and pcl filenames must be provided")
		cmd.Usage()
		return
	}

	if _, err := os.Stat(input); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", input)
		cmd.Usage()
		return
	}

	if _, err := os.Stat(pcl); os.IsNotExist(err) {
		fmt.Printf("No such file or directory: %s\n", pcl)
		cmd.Usage()
		return
	}

	if _, err := os.Stat(output); err == nil {
		fmt.Printf("%s exists; overwriting...\n", output)
	}

	readerType, _ := utils.InferReaderFromExt(input)
	writerType, _ := utils.InferWriterFromExt(output)

	vg := `--filters.pclblock.filename=` + pcl

	utils.RunPdalOmni(input, output,
		"-r", readerType, "-w", writerType, "--filter", "filters.pclblock",
		vg, "-v10", "--debug")

	if view {
		utils.OpenData(output)
	}
}
