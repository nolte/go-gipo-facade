package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/nolte/go-gipo-facade/pkg"
	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "GPIOs Command Line",
	Long:  `CLI Helper for PI Gpio ussage`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		log.Info("Starting")
		pkg.ControllGIPO()
	},
}
