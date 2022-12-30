package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

const encryptCost = bcrypt.DefaultCost

var cmdPassword = &cobra.Command{
	Use:   "password",
	Short: "Password related commands",
	Run: func(cmd *cobra.Command, args []string) {
		// default command: print usage
		cmd.Usage()
	},
}

var cmdPasswordHash = &cobra.Command{
	Use:   "hash",
	Short: "Hash a password",
	Args:  cobra.ExactArgs(1),
	Run:   runPasswordHash,
}

var cmdPasswordCompare = &cobra.Command{
	Use:   "compare",
	Short: "Compare a password against a hashed password",
	Args:  cobra.ExactArgs(2),
	Run:   runPasswordCompare,
}

func init() {
	cmdPassword.AddCommand(cmdPasswordHash)
	cmdPassword.AddCommand(cmdPasswordCompare)

	rootCmd.AddCommand(cmdPassword)
}

// Hash a password
func runPasswordHash(cmd *cobra.Command, args []string) {
	password := args[0]
	hash, err := bcrypt.GenerateFromPassword([]byte(password), encryptCost)
	if err != nil {
		log.Fatalf("Fail to hash password: %v", err)
	}
	fmt.Println(string(hash))
}

// Compare a password against a hashed password
func runPasswordCompare(cmd *cobra.Command, args []string) {
	password := args[0]
	hashed := args[1]
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		log.Fatalf("invalid password: %v", err)
	}
	fmt.Println("Password matched.")
}
