package main

import (
	"fmt"
	"os"

	"github.com/arschles/helium/helium-cli/actions"
	log "github.com/arschles/helium/pkg/log/human"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "he",
		Short:        "Your one stop shop for building and deploying Kubernetes apps",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: version.
			// see https://github.com/Azure/service-catalog-cli/blob/master/cmd/svcat/main.go#L34
			fmt.Print(cmd.UsageString())
			return nil
		},
	}
	cmd.AddCommand(actions.Run())
	cmd.AddCommand(actions.RunServer())
	flags := cmd.PersistentFlags()
	var debug bool
	flags.BoolVarP(&debug, "debug", "d", false, "Turn on debug logging")
	flags.Parse(os.Args)
	log.IsDebugging = debug
	// TODO:
	// 	cli.StringFlag{
	// 		Name:  actions.FlagConfigDir,
	// 		Usage: "The location of the config directory",
	// 	},

	// app.Before = func(c *cli.Context) error {
	// 	log.IsDebugging = c.Bool(actions.FlagDebug)
	// 	return nil
	// }

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
