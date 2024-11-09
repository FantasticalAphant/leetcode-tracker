package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all problems",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(problemStore.Problems) == 0 {
			fmt.Println("No problems found.")
			return nil
		}

		fmt.Println("Problems:")
		for _, problem := range problemStore.Problems {
			status := " "
			if problem.Completed {
				status = "✓"
			}
			fmt.Printf("[%s] %d (Updated: %v)\n", status, problem.ID, problem.Modified)
		}
		return nil
	},
}
