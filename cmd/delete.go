package cmd

import (
	"github.com/JackalLabs/jlaunch/core"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func DeleteCMD() *cobra.Command {
	r := cobra.Command{
		Use:   "delete [folder]",
		Short: "delete removes a folder from the jackal network",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			operatingRoot := args[0]

			log.Info().Msgf("Removing `%s` !", operatingRoot)

			err := core.DeleteFolder(operatingRoot)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return &r
}
