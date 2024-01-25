package cmd

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

func RootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "jlaunch",
		Short: "jlaunch launches an entire project into the cloud",

		RunE: func(cmd *cobra.Command, args []string) error {

			log.Info().Msg("Launching files to the cloud!")

			return nil
		},
	}

	root.AddCommand(VersionCmd())

	return root
}

func Execute(root *cobra.Command) {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
