package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"go-swagger-example/logger"
)

// this server program version string
const name = "example"
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
	fmt.Printf("%v v%v\n", name, version)
}

func logVersion(log logger.Logger) {
	log.Infof("%v v%v", name, version)
}
