package main

import (
	"github.com/JackalLabs/jlaunch/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra/doc"
	"os"
)

func main() {
	root := cmd.RootCmd()

	isProdEnv := os.Getenv("PROD")

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
	if isProdEnv == "false" {
		os.RemoveAll("./doc/")
		log.Logger = log.Level(zerolog.DebugLevel)
		_ = os.Mkdir("./doc/", os.ModePerm)
		err := doc.GenMarkdownTree(root, "./doc/")
		if err != nil {
			log.Error().Err(err)
			return
		}
	} else {
		log.Logger = log.Level(zerolog.InfoLevel)
	}

	cmd.Execute(root)
}
