/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/cage1016/alfred-mergeimgs/alfred"
	"github.com/spf13/cobra"
)

// mgfCmd represents the mgf command
var mgfCmd = &cobra.Command{
	Use:   "mgf",
	Short: "List all the mgf folders",
	Run:   runMgfCmd,
}

func runMgfCmd(cmd *cobra.Command, args []string) {
	data, _ := alfred.LoadOngoingSources(wf)

	log.Println("data", data)

	for name, path := range data {
		wi := wf.NewItem(name).
			Subtitle(fmt.Sprintf("⌘ ,↩ Select files to Merge from '%s'", path)).
			Valid(true).
			Arg(path)

		wi.Cmd().
			Subtitle("↩ Enter Action menu to Add / Remove source folder").
			Valid(true).
			Arg("")
	}

	if len(data) == 0 {
		wf.NewItem("No any source available").
			Subtitle("↩ Add more source folder to configuration").
			Valid(true).
			Arg("")
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(mgfCmd)
}
