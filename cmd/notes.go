package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	notesCmd.Flags().BoolVarP(&add, "add", "a", false, "add notes")
	rootCmd.AddCommand(notesCmd)
}

var add bool

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Notes for each question",
	Args:  cobra.ExactArgs(1), // not sure what the command is going to look like
	RunE: func(cmd *cobra.Command, args []string) error {
		switch {
		case add:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter notes: ")

			text, _ := reader.ReadString('\n')
			if err := problemStore.AddNote(text, args[0]); err != nil {
				return err
			}
		default:
			notes, err := problemStore.ShowNotes(args[0])
			if err != nil {
				return err
			}
			fmt.Println(notes)
		}
		return nil
	},
}
