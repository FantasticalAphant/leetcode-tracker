package cmd

import (
	"fmt"
	"heatcold/internal/leetcode"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get more information about a LeetCode problem",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		name, err := leetcode.GetQuestionNameByID(id)
		level, err := leetcode.GetQuestionDifficultyByID(id)
		if err != nil {
			return err
		}
		fmt.Println(name)
		fmt.Println(level)
		return nil
	},
}
