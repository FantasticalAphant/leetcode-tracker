package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().BoolVarP(&long, "long", "l", false, "show more information")
	rootCmd.AddCommand(listCmd)
}

var long bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all problems",
	// TODO: allow user to specify questions to print out
	Args: cobra.MaximumNArgs(1),
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
			fmt.Printf("[%s] %d", status, problem.ID)
			if long {
				fmt.Printf(" (Updated): %v\tNotes: %v", problem.Modified, problem.Notes)
			}
			fmt.Println()
		}
		return nil
	},
}
