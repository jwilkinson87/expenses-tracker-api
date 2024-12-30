package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "migrateHow",
	Short: "migration is a tool for running database migrations",
	Long:  "migration is at tool for running all SQL scripts within a specified directory.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sqlDirectory := args[0]
		fmt.Println(sqlDirectory)
		if _, err := os.Stat(sqlDirectory); err != nil {
			fmt.Fprintln(os.Stderr, "SQL directory cannot be accessed. Does it exist?")
			os.Exit(1)
		}

		handler := NewMigrationHandler(sqlDirectory)
		handler.Execute()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
