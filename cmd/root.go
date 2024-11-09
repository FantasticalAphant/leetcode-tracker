package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"heatcold/internal/store"
)

var problemStore *store.ProblemStore

var rootCmd = &cobra.Command{
	Use:   "yeetcode",
	Short: "Track your leetcode progress",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		problemStore, err = store.New()
		if err != nil {
			return err
		}
		return problemStore.Load()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You are calling this without any arguments")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
