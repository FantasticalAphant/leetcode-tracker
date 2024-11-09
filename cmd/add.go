package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new completed leetcode problem",
	Long:  `Add a new completed leetcode problem`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := problemStore.AddProblem(args[0]); err != nil {
			return err
		}
		fmt.Println("Problem added successfully")
		return nil
	},
}
