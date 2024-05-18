package cmd

import (
	"github.com/JackalLabs/jlaunch/core"
	"github.com/JackalLabs/jutils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"path"
	"path/filepath"
)

func LaunchCmd() *cobra.Command {
	r := cobra.Command{
		Use:   "launch [folder]",
		Short: "launch saves an entire folder to the jackal network",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			p, err := filepath.Abs(args[0])
			if err != nil {
				return err
			}

			log.Info().Msgf("Launching `%s` !", p)

			root := jutils.LoadEnvVarOrFallback("ROOT", "launch")
			operatingRoot := "s/" + root

			err = core.SaveFolder(p, operatingRoot)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return &r
}

func LaunchFileCmd() *cobra.Command {
	r := cobra.Command{
		Use:   "launch-file [file] [folder]",
		Short: "launch saves a file to the jackal network",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			file := args[0]
			folder := args[1]

			log.Info().Msgf("Launching `%s` !", file)

			root := jutils.LoadEnvVarOrFallback("ROOT", "launch")
			operatingRoot := "s/" + root

			operatingRoot = path.Join(operatingRoot, folder)

			err := core.SaveFile(file, operatingRoot)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return &r
}
