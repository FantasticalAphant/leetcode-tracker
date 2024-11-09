package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a leetcode problem from the list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := problemStore.RemoveProblem(args[0]); err != nil {
			return err
		}
		fmt.Println("Problem successfully deleted")
		return nil
	},
}
