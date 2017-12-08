package main

import (
	"fmt"
	"os"

	"github.com/kubehelium/helium/cli/actions"
	// log "github.com/kubehelium/helium/pkg/log/human"
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
	// TODO:
	//
	// app.Flags = []cli.Flag{
	// 	cli.BoolFlag{
	// 		Name:  actions.FlagDebug,
	// 		Usage: "Enable verbose debugging output",
	// 	},
	// 	cli.StringFlag{
	// 		Name:  actions.FlagConfigDir,
	// 		Usage: "The location of the config directory",
	// 	},
	// }

	// app.Before = func(c *cli.Context) error {
	// 	log.IsDebugging = c.Bool(actions.FlagDebug)
	// 	return nil
	// }

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
