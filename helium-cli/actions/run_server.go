package actions

import (
	"os"

	log "github.com/arschles/helium/pkg/log/human"
	"github.com/arschles/helium/pkg/runtime"
	"github.com/spf13/cobra"
)

type serverRunner struct {
	port int
}

func (s *serverRunner) run(cmd *cobra.Command, args []string) error {
	log.Info("Starting the runtime server on port %d", s.port)
	srv := runtime.NewServer(s.port)
	return srv.Run()
}

func RunServer() *cobra.Command {
	runner := &serverRunner{}
	cmd := &cobra.Command{
		Use:          "run-server",
		Short:        "Start the runtime server locally",
		SilenceUsage: true,
		RunE:         runner.run,
	}
	flags := cmd.Flags()
	flags.IntVarP(&runner.port, "port", "p", 8080, "The port to run the server on")
	flags.Parse(os.Args)

	return cmd
}
