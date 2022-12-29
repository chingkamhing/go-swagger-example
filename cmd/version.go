package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// this server program version string
const version = "0.0.1"

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(cmdVersion)
}

func runVersion(cmd *cobra.Command, args []string) {
	// read version file
	fmt.Printf("Example Server v%v\n", version)
}
