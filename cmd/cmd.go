package cmd

import (
	"time"

	"github.com/spf13/cobra"

	"go-swagger-example/logger"
)

// default settings
const defaultConfigFile string = "config.yaml" // default config yaml file

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Web Server",
	Run: func(cmd *cobra.Command, args []string) {
		// default command
		cmd.Usage()
	},
}

func init() {
	rootCmd.PersistentFlags().String("config", defaultConfigFile, "config file")
	rootCmd.PersistentFlags().String("database.accounts", "", "DBAccounts web config string (note: this has higher precedence than other database settings)")
	rootCmd.PersistentFlags().String("database.driver", "mysql", "Database driver of: mysql")
	rootCmd.PersistentFlags().String("database.host", "localhost", "Database host string")
	rootCmd.PersistentFlags().Int("database.port", 1433, "Database port")
	rootCmd.PersistentFlags().String("database.name", "ExampleDB", "Database name")
	rootCmd.PersistentFlags().String("database.user", "", "Database user name")
	rootCmd.PersistentFlags().String("database.password", "", "Database password")
	rootCmd.PersistentFlags().Int("database.retry", 10, "Database connection retry count")
	rootCmd.PersistentFlags().Duration("database.interval", time.Duration(3*time.Second), "Database connection retry interval in second")

	// set --version flag
	rootCmd.Version = version
	rootCmd.InitDefaultVersionFlag()
}

// Execute init cli commands, flags and read configuration
func Execute() {
	// run root command
	err := rootCmd.Execute()
	if err != nil {
		logger.Fatalln(err)
	}
}
