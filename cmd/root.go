package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	tui2 "heatcold/internal/tui"
	"os"

	"github.com/spf13/cobra"
	"heatcold/internal/store"
)

func init() {
	rootCmd.Flags().BoolVar(&tui, "tui", false, "use tui")
}

var problemStore *store.ProblemStore
var tui bool

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
		if tui {
			p := tea.NewProgram(tui2.InitialModel(), tea.WithAltScreen())
			if _, err := p.Run(); err != nil {
				os.Exit(1)
			}
		} else {
			fmt.Println("You are calling this without any arguments")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
