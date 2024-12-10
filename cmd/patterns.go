package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"heatcold/internal/leetcode"
)

func init() {
	rootCmd.AddCommand(patternsCmd)
}

var patternsCmd = &cobra.Command{
	Use:   "patterns",
	Short: "LeetCode question patterns",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		patterns, err := leetcode.GetCodingPatterns()
		if err != nil {
			return err
		}
		fmt.Println()
		for i, pattern := range patterns {
			fmt.Printf("%d: %s\n", i+1, pattern)
		}

		return nil
	},
}
