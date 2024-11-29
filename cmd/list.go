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
			fmt.Printf("[%s] %-4d", status, problem.ID)
			if long {
				fmt.Printf(" | Updated: %v | Notes: %v", problem.Modified.Format("01/02/2006 @ 15:04:05"), problem.Notes)
			}
			fmt.Println()
		}
		return nil
	},
}
