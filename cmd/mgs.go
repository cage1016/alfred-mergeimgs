/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-mergeimgs/alfred"
)

// mgsCmd represents the ms command
var mgsCmd = &cobra.Command{
	Use:   "mgs",
	Short: "Merge Image settings",
	Run:   runtMgsCmd,
}

func runtMgsCmd(cmd *cobra.Command, args []string) {
	data, _ := alfred.LoadOngoingSources(wf)

	wf.NewItem("BACK").
		Subtitle("Back to list").
		Valid(true).
		Var("action", "back")

	wf.NewItem("ADD").
		Subtitle("Add more source folder to configuration").
		Valid(true).
		Arg("").
		Var("action", "add")

	for name, path := range data {
		wf.NewItem(fmt.Sprintf("REMOVE %s", name)).
			Subtitle(fmt.Sprintf("remove '%s' from configuration", path)).
			Valid(true).
			Arg(path).
			Var("action", "remove")
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(mgsCmd)
}
