package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update status for the leetcode problem",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := problemStore.UpdateProblem(args[0]); err != nil {
			return err
		}
		fmt.Println("Problem updated successfully")
		return nil
	},
}
