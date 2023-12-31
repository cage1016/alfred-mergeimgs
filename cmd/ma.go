/*
Copyright © 2022 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-mergeimgs/alfred"
)

// maCmd represents the ma command
var maCmd = &cobra.Command{
	Use:   "ma",
	Short: "Manage Merge Images Sources",
	Run:   runMaCmd,
}

func runMaCmd(cmd *cobra.Command, args []string) {
	data, _ := alfred.LoadOngoingSources(wf)
	action, _ := cmd.Flags().GetString("action")
	folder := filepath.Base(args[0])

	switch action {
	case "add":
		data[folder] = args[0]
		alfred.StoreOngoingSources(wf, data)
	case "remove":
		delete(data, folder)
		alfred.StoreOngoingSources(wf, data)
	default:
		logrus.Info("Do nothing")
	}
}

func init() {
	rootCmd.AddCommand(maCmd)
	maCmd.PersistentFlags().StringP("action", "a", "", "type of Action")
}
