package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	Version = "v0.0.0"
	Commit  = ""
)

func VersionCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "version",
		Short: "version prints the version of the jlaunch binary",

		RunE: func(cmd *cobra.Command, args []string) error {

			log.Info().Msgf("Version: %s", Version)
			log.Info().Msgf("Commit: %s", Commit)
			return nil
		},
	}

	return root
}
