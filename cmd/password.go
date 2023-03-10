package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"

	"go-swagger-example/gen/models"
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

var cmdPasswordUser = &cobra.Command{
	Use:   "user",
	Short: "JSON encode user login info",
	Args:  cobra.ExactArgs(0),
	Run:   runPasswordUser,
}

func init() {
	cmdPassword.AddCommand(cmdPasswordHash)
	cmdPassword.AddCommand(cmdPasswordCompare)
	cmdPassword.AddCommand(cmdPasswordUser)

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

// JSON encode user login info
func runPasswordUser(cmd *cobra.Command, args []string) {
	user := models.UserLogin{
		UserID:       0,
		LoginName:    "user-1001",
		Email:        "user-1001@email.com",
		PasswordHash: "password123456",
	}
	fmt.Printf("user: %v\n", user)
	userEncoded, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("json marshal error: %v", err)
	}
	fmt.Printf("userEncoded: %v\n", string(userEncoded))
}
