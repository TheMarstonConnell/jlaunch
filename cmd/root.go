package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func RootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "jlaunch",
		Short: "jlaunch launches an entire project into the cloud",
	}

	root.AddCommand(VersionCmd(), LaunchCmd(), LaunchFileCmd(), DeleteCMD())

	return root
}

func Execute(root *cobra.Command) {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
