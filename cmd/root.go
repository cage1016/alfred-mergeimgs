/*
Copyright © 2023 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"errors"
	"log"
	"os"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-mergeimgs/alfred"
)

var (
	repo = "cage1016/alfred-mergeimgs"
	wf   *aw.Workflow
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Merge Images",
	Short: "Merge the images into a singular horizontal or vertical image",
	Run: func(cmd *cobra.Command, args []string) {
		wf.SendFeedback()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	wf.Run(func() {
		if err := rootCmd.Execute(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	})
}

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
	wf.Args() // magic for "workflow:update"

	// Create init dir if it doesn't exist
	if _, err := os.Stat(alfred.GetDestDir(wf)); errors.Is(err, os.ErrNotExist) {
		os.Mkdir(alfred.GetDestDir(wf), 0755)
	}

	if alfred.GetDebug(wf) {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
