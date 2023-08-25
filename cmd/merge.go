/*
Copyright © 2023 KAI CHU CHUNG <cage.chung@gmail.com>
*/
package cmd

import (
	"os/exec"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/cage1016/alfred-mergeimgs/alfred"
	"github.com/cage1016/alfred-mergeimgs/lib"
)

var (
	av = aw.NewArgVars()
)

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge the images into a singular horizontal or vertical image",
	Run:   runMergeCmd,
}

var appendCmd = map[string]string{
	"horizontal": "+smush",
	"vertical":   "-smush",
}

func ErrorHandle(err error) {
	logrus.Errorf("ErrorHandle: %v", err)
	av.Var("err", err.Error())
	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args to Alfred: %v", err)
	}
}

func runMergeCmd(cmd *cobra.Command, args []string) {
	o, _ := cmd.Flags().GetString("output")
	m, _ := cmd.Flags().GetString("mode")
	l, _ := cmd.Flags().GetInt("length")

	// check fd command
	if err := lib.CheckCommand("fd"); err != nil {
		ErrorHandle(errors.Wrap(err, "failed to check fd command"))
		return
	}

	// check imagemagick command
	if err := lib.CheckCommand("convert"); err != nil {
		ErrorHandle(errors.Wrap(err, "failed to check imagemagick command"))
		return
	}

	var output = filepath.Join(o, lib.RandomImageFilename(l))
	var ok bool
	var ap string

	if ap, ok = appendCmd[m]; !ok {
		ErrorHandle(errors.Errorf("invalid mode: %s. Only support horizontal or vertical", m))
		return
	}

	cmds := []string{ap, alfred.GetOffset(wf)}

	// gravity
	cmds = append(cmds, "-gravity", alfred.GetGravity(wf))

	// background
	if alfred.GetBackground(wf) != "" {
		cmds = append(cmds, "-background", alfred.GetBackground(wf))
	}

	cmds = append(cmds, args...)
	cmds = append(cmds, output)
	err := runCmd("convert", cmds...)

	logrus.Debugf("cmds: convert %v", strings.Join(cmds, " "))

	if err != nil {
		ErrorHandle(errors.Wrap(err, "failed to merge images"))
		return
	} else {
		av.Var("err", "")
		av.Arg(output)
	}

	if err := av.Send(); err != nil {
		wf.Fatalf("failed to send args: %v", err)
	}
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().StringP("output", "o", "", "output folder")
	mergeCmd.Flags().StringP("mode", "m", "", "merge mode: horizontal or vertical")
	mergeCmd.Flags().IntP("length", "l", 10, "random file name length")
}

func runCmd(cmd string, args ...string) error {
	if _, err := exec.Command(cmd, args...).CombinedOutput(); err != nil {
		return errors.Wrap(err, "failed to run command")
	}
	return nil
}
