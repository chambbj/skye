package commands

//
// import (
// 	"fmt"
// 	"math"
// 	"os"
//
// 	"github.com/chambbj/skye/utils"
// 	"github.com/spf13/cobra"
// )
//
// var splitCmd = &cobra.Command{
// 	Use:   "split",
// 	Short: "Split input files",
// 	Long:  "",
// 	Run:   RunSplit,
// }
//
// var length, origin_x, origin_y float64
// var capacity uint32
//
// func init() {
// 	splitCmd.Flags().StringVarP(&input, "input", "i", "", "Input filename")
// 	splitCmd.Flags().StringVarP(&output, "output", "o", "", "Output filename")
// 	splitCmd.Flags().Uint32VarP(&capacity, "capacity", "", 0, "Point capacity of chipper cells")
// 	splitCmd.Flags().Float64VarP(&length, "length", "", 0.0, "Edge length for splitter cells")
// 	splitCmd.Flags().Float64VarP(&origin_x, "origin_x", "", math.NaN(), "Origin in X axis for splitter cells")
// 	splitCmd.Flags().Float64VarP(&origin_y, "origin_y", "", math.NaN(), "Origin in Y axis for splitter cells")
// }
//
// func RunSplit(cmd *cobra.Command, args []string) {
// 	if input == "" || output == "" {
// 		fmt.Println("input and output filenames must be provided")
// 		cmd.Usage()
// 		return
// 	}
//
// 	if _, err := os.Stat(input); os.IsNotExist(err) {
// 		fmt.Printf("No such file or directory: %s\n", input)
// 		cmd.Usage()
// 		return
// 	}
//
// 	if _, err := os.Stat(output); err == nil {
// 		fmt.Printf("%s exists; overwriting...\n", output)
// 	}
//
// 	if length && capacity {
// 		fmt.Println("Can't specify for length and capacity.")
// 		cmd.Usage()
// 		return
// 	}
//
// 	if !length && !capacity {
// 		capacity = 100000
// 	}
//
// 	// if output file is a directory, append the input file name
//
// 	readerType, _ := utils.InferReaderFromExt(input)
// 	writerType, _ := utils.InferWriterFromExt(output)
//
// 	// it will be difficult to implement this here because the split kernel
// 	// actually creates multiple writers based off the single reader
//
// 	if length {
// 		// ignoring x/y origin right now
// 		lengthOpt := "--filters.splitter.length=" + length
// 		xOriginOpt := "--filters.splitter.origin_x=" + origin_x
// 		yOriginOpt := "--filters.splitter.origin_y=" + origin_y
// 		utils.RunPdalOmni(input, output,
// 			"-r", readerType, "-w", writerType, "--filter", "filters.splitter",
// 			lengthOpt, xOriginOpt, yOriginOpt, "-v10", "--debug")
// 	} else {
// 		capacityOpt := "--filters.chipper.capacity=" + capacity
// 		utils.RunPdalOmni(input, output,
// 			"-r", readerType, "-w", writerType, "--filter", "filters.chipper",
// 			"-v10", "--debug")
// 	}
//
//  if view {
// 	  utils.OpenData(output)
//  }
// }
