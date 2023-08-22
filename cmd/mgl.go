/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-mergeimgs/alfred"
	"github.com/cage1016/alfred-mergeimgs/lib"
)

// mglCmd represents the mgl command
var mglCmd = &cobra.Command{
	Use:   "mgl",
	Short: "Merge Image files list",
	Run:   runMglCmd,
}

func runMglCmd(cmd *cobra.Command, args []string) {
	source, _ := cmd.Flags().GetString("target")

	var err error
	ranges := &[]lib.Range{}
	query := args[0]
	if lib.IsPageRangeValid(args[0]) {
		page := lib.PageRangeRegex.FindAllString(args[0], -1)
		query = strings.TrimSpace(strings.Replace(args[0], page[0], "", -1))
		ranges, err = lib.ParseRangeNumber(strings.Replace(page[0], "#", "", -1), alfred.GetMaxQueryResults(wf))
		if err != nil {
			logrus.Errorf("failed to parse page range: %v", err)
			return
		}
	}

	documents := lib.FdExecute(lib.FdConfig{
		Source:          source,
		Arg:             query,
		Type:            "-tf\n-tl",
		Extension:       alfred.GetQueryExtension(wf),
		MaxDepth:        "-d 1",
		MaxQueryResults: alfred.GetMaxQueryResults(wf),
	})

	nTargets := []string{}
	for _, r := range *ranges {
		if r.Start > 0 {
			if r.Start == r.End {
				nTargets = append(nTargets, documents[r.Start-1])
			} else {
				nTargets = append(nTargets, documents[r.Start-1:r.End]...)
			}
		}
	}

	if len(documents) > 0 {
		if !lib.IsFdError(documents[0]) {
			for i, doc := range documents {
				if strings.TrimSpace(doc) == "" {
					continue
				}
				i += 1

				if len(*ranges) == 0 {
					wi := wf.NewItem(fmt.Sprintf("%d - %s", i, filepath.Base(doc))).
						Subtitle("⌥, Press return to merge horizontal recent files up to this file.").
						Quicklook(doc).
						Valid(true).
						Arg(documents[:i]...).
						ActionForType("file", documents[:i]...).
						Icon(&aw.Icon{
							Value: doc,
						})

					wi.Alt().
						Subtitle("Press return to merge vertical recent files up to this file.").
						Valid(true).
						Arg(documents[:i]...)

				} else {
					var wi *aw.Item
					if lib.Ranges(*ranges).IsInRange(i) {
						wi = wf.NewItem(fmt.Sprintf("✅ %d - %s", i, filepath.Base(doc))).
							Subtitle("⌥, Press return to merge horizontal select files in order.")
					} else {
						wi = wf.NewItem(fmt.Sprintf("%d - %s", i, filepath.Base(doc)))
					}

					wi.Quicklook(doc).
						Valid(true).
						Arg(nTargets...).
						ActionForType("file", nTargets...).
						Icon(&aw.Icon{
							Value: doc,
						})

					wi.Alt().
						Subtitle("Press return to merge vertical select files in order.").
						Valid(true).
						Arg(nTargets...)
				}
			}
		} else {
			for _, doc := range documents {
				wf.NewItem(doc).
					Subtitle("Press return to visit fd document").
					Valid(true).
					Var("action", "open").
					Var("help", "https://github.com/sharkdp/fd")
			}
		}
	}

	wf.SendFeedback()
}

func init() {
	rootCmd.AddCommand(mglCmd)
	mglCmd.Flags().StringP("target", "t", "", "target folder")
}
