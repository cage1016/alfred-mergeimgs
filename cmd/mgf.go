/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-mergeimgs/alfred"
	"github.com/cage1016/alfred-mergeimgs/lib"
)

// mgfCmd represents the mgf command
var mgfCmd = &cobra.Command{
	Use:   "mgf",
	Short: "List all the mgf folders",
	Run:   runMgfCmd,
}

func runMgfCmd(cmd *cobra.Command, args []string) {
	// check fd command
	if err := lib.CheckCommand("fd"); err != nil {
		wf.NewItem("fd not found").
			Subtitle("Press return to visit 'https://github.com/sharkdp/fd' install fd first").
			Valid(true).
			Arg("https://github.com/sharkdp/fd")

		wf.SendFeedback()
		return
	}

	// check imagemagick command
	if err := lib.CheckCommand("convert"); err != nil {
		wf.NewItem("imagemagick not found").
			Subtitle("Press return to visit 'https://github.com/ImageMagick/ImageMagick' install imagemagick first").
			Valid(true).
			Arg("https://github.com/ImageMagick/ImageMagick")

		wf.SendFeedback()
		return
	}

	data, _ := alfred.LoadOngoingSources(wf)
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
