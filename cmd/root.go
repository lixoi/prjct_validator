package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "prjctvalid",
	Short: "project validator",
}

func init() {
	RootCmd.AddCommand(StartValidator)
}
