package cmd

import (
	"fmt"
	"heatcold/internal/leetcode"

	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().BoolVarP(&long, "long", "l", false, "show more information")
	listCmd.Flags().BoolVarP(&info, "info", "i", false, "show api information")
	listCmd.MarkFlagsMutuallyExclusive("long", "info")
	rootCmd.AddCommand(listCmd)
}

var long bool
var info bool

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

		var questionInfo map[int]leetcode.QuestionInformation

		if info {
			var err error
			questionInfo, err = leetcode.GetQuestionRangeInformation(problemStore.GetProblemsSorted()...)
			if err != nil {
				return err
			}
		}

		for _, key := range problemStore.GetProblemsSorted() {
			problem := problemStore.Problems[key]

			status := " "
			if problem.Completed {
				status = "✓"
			}
			fmt.Printf("[%s] %-5d", status, problem.ID)
			if info {
				fmt.Printf("| %s (%s)", questionInfo[problem.ID].Name, questionInfo[problem.ID].Difficulty)
			}
			if long {
				fmt.Printf("| Updated: %v | Notes: %v", problem.Modified.Format("01/02/2006 @ 15:04:05"), problem.Notes)
			}
			fmt.Println()
		}

		return nil
	},
}
