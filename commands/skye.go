package commands

import (
	"fmt"

	"github.com/chambbj/skye/utils"
	"github.com/spf13/cobra"
)

var input, output string
var reader, writer string
var view bool

// SkyeCmd is Skye's root command. Every other command attached to SkyeCmd is a
// child command to it.
var SkyeCmd = &cobra.Command{
	Use:   "skye",
	Short: "skye short",
	Long: `skye is the main command.

Skye is a wrapper for PDAL.`,
}
var skyeCmdV *cobra.Command

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Skye",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Skye v0.1 -- HEAD")
	},
}

var driversCmd = &cobra.Command{
	Use:   "drivers",
	Short: "Print PDAL's available drivers",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		utils.RunPdal("--drivers")
	},
}

// Execute adds all child commands to the root command SkyeCmd and sets flags
// appropriately.
func Execute() {
	AddCommands()
	SkyeCmd.Execute()
}

// AddCommands adds child commands to the root SkyeCmd.
func AddCommands() {
	SkyeCmd.AddCommand(colorizeCmd)
	SkyeCmd.AddCommand(convertCmd)
	SkyeCmd.AddCommand(cropCmd)
	SkyeCmd.AddCommand(driversCmd)
	SkyeCmd.AddCommand(groundCmd)
	SkyeCmd.AddCommand(infoCmd)
	SkyeCmd.AddCommand(outlierCmd)
	SkyeCmd.AddCommand(pclCmd)
	SkyeCmd.AddCommand(pipelineCmd)
	SkyeCmd.AddCommand(sortCmd)
	SkyeCmd.AddCommand(thinCmd)
	SkyeCmd.AddCommand(versionCmd)
	SkyeCmd.AddCommand(voxelGridCmd)
}

func init() {
	SkyeCmd.PersistentFlags().BoolVarP(&view, "view", "v", false,
		"View output using default application")

	skyeCmdV = SkyeCmd
}
