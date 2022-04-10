package pgk

import (
	"CliTaskManager/cmd/Todo/handler"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

// configCmd represents the config command
var addNotes = &cobra.Command{
	Use:   "add",
	Short: "Istenilen yerin sicakligini verir ",
	Run: func(cmd *cobra.Command, args []string) {
		handler.CreateNewNote(strings.Join(args, ""))
	},
}
var listNote = &cobra.Command{
	Use:   "list",
	Short: "notlar listelenir",
	Run: func(cmd *cobra.Command, args []string) {
		handler.GetAllNotes()
	},
}
var deleteNote = &cobra.Command{
	Use:   "delete",
	Short: "notlar silinir",
	Run: func(cmd *cobra.Command, args []string) {
		index, _ := strconv.Atoi(strings.Join(args, ""))
		handler.DeleteNote(index)
	},
}

func init() {
	RootCmd.AddCommand(addNotes)
	RootCmd.AddCommand(listNote)
	RootCmd.AddCommand(deleteNote)
}
