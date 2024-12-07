package cmd

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"heatcold/internal/leetcode"
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

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

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
			info, err := leetcode.GetQuestionInformation(1)
			if err != nil {
				fmt.Println("Invalid")
				os.Exit(1)
			}
			m := tui2.Model{List: list.New([]list.Item{item{title: info["name"], desc: info["difficulty"]}}, list.NewDefaultDelegate(), 0, 0)}
			m.List.Title = "List of Questions"
			p := tea.NewProgram(m, tea.WithAltScreen())
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
