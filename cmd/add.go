package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	addCmd.Flags().BoolVarP(&completed, "completed", "c", false, "mark newly added question as complete")
	rootCmd.AddCommand(addCmd)
}

var completed bool

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new completed leetcode problem",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := problemStore.AddProblem(completed, args); err != nil {
			return err
		}
		fmt.Println("Problem added successfully")
		return nil
	},
}
