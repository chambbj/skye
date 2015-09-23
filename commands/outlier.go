package commands

import "github.com/spf13/cobra"

var outlierCmd = &cobra.Command{
	Use:   "outlier",
	Short: "Remove outliers",
	Long:  "",
}

var method string

func init() {
	outlierCmd.AddCommand(radiusCmd)
	outlierCmd.AddCommand(statisticalCmd)
}
