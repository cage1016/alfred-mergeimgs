/*
Copyright © 2023 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"os/exec"
	"path/filepath"

	aw "github.com/deanishe/awgo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-mergeimgs/lib"
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge the images into a singular horizontal or vertical image",
	Run:   runMergeCmd,
}

var appendCmd = map[string]string{
	"horizontal": "+append",
	"vertical":   "-append",
}

func runMergeCmd(cmd *cobra.Command, args []string) {
	o, _ := cmd.Flags().GetString("output")
	m, _ := cmd.Flags().GetString("mode")

	var output = filepath.Join(o, lib.RandomImageFilename(10))
	var ok bool
	var ap string

	if ap, ok = appendCmd[m]; !ok {
		wf.Fatalf("invalid mode: %s", m)
	}

	cmds := []string{ap}
	cmds = append(cmds, args...)
	cmds = append(cmds, output)
	err := runCmd("convert", cmds...)

	av := aw.NewArgVars()
	av.Arg(output)
	if err != nil {
		logrus.Printf("failed to merge images: %v", err)
		av.Var("err", err.Error())
	} else {
		av.Var("err", "")
	}
	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().StringP("output", "o", "", "output folder")
	mergeCmd.Flags().StringP("mode", "m", "", "merge mode: horizontal or vertical")
}

func runCmd(cmd string, args ...string) error {
	if _, err := exec.Command(cmd, args...).CombinedOutput(); err != nil {
		return errors.Wrap(err, "failed to run command")
	}
	return nil
}
