package actions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	// "github.com/kubehelium/helium/cli/util/docker"
	log "github.com/kubehelium/helium/pkg/log/human"
	"github.com/kubehelium/helium/pkg/runtime"
)

type runner struct {
	port  int
	heDir string
}

func (r *runner) run(c *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: he run <event>")
	}
	event := args[0]
	// event := c.String("target")
	// dockerClient := docker.ClientOrDie()
	// TODO: make configurable
	const port = 8080
	log.Info("build")
	srv := runtime.NewServer(port)
	srvErrCh := make(chan error)
	go func() {
		log.Debug("Starting the runtime server")
		srvErrCh <- srv.Run()
	}()
	// TODO: ping the server until it's up

	imgStr := "helium-runner:v0.0.1"
	// img, err := docker.ParseImageFromName(imgStr)
	// if err != nil {
	// log.Err("error parsing docker image %s (%s)", imgStr, err)
	// os.Exit(1)
	// }

	// rmContainerCh := make(chan func())
	// stdOutCh := make(chan docker.Log)
	// stdErrCh := make(chan docker.Log)
	// exitCodeCh := make(chan int)
	// errCh := make(chan error)

	// TODO: make configurable
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working dir")
	}
	absHeDir := filepath.Join(wd, r.heDir)
	log.Debug("Using He build directory %s", r.heDir)
	// TODO: also start up the helium runtime server, and link them together
	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-v",
		fmt.Sprintf("%s:/helium", absHeDir),
		"--net",
		"host",
		"-e",
		"HELIUM_SCRIPTS_DIR=/helium",
		"-e",
		"HELIUM_TARGET_DIR=/home/src/dist/src",
		"-e",
		fmt.Sprintf("HELIUM_EVENT_TYPE=%s", event),
		"-e",
		"HELIUM_EVENT_PROVIDER=cli",
		"-e",
		fmt.Sprintf("HELIUM_URL=%d", port),
		"-p",
		fmt.Sprintf("127.0.0.1:%d:%d", port, port),
		// TODO: HELIUM_EVENT_METADATA?
		imgStr,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debug(strings.Join(cmd.Args, " "))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running command (%s)", err)
	}
	// go docker.Run(
	// 	dockerClient,
	// 	img,
	// 	"helium-runner",

	// 	docker.ContainerGopath(cfg.Build.Gopath, paths.PackageName),
	// 	fmt.Sprintf("go build -o %s .", binaryName),
	// 	cfg.Build.Env,
	// 	rmContainerCh,
	// 	stdOutCh,
	// 	stdErrCh,
	// 	exitCodeCh,
	// 	errCh,
	// )

	// for {
	// 	select {
	// 	case rmContainerFn := <-rmContainerCh:
	// 		defer rmContainerFn()
	// 	case l := <-stdOutCh:
	// 		log.Info("%s", l)
	// 	case l := <-stdErrCh:
	// 		log.Warn("%s", l)
	// 	case err := <-errCh:
	// 		log.Err("%s", err)
	// 		return
	// 	case i := <-exitCodeCh:
	// 		log.Info("exited with code %d", i)
	// 		return
	// 	}
	// }
	return nil
}

// Run is the CLI handler for 'he run'
func Run() *cobra.Command {
	runner := &runner{}
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Execute an individual event on the helium build script",
		RunE:  runner.run,
	}
	flags := cmd.Flags()
	flags.IntVarP(
		&runner.port,
		"runtime-port",
		"p",
		8080,
		"The port on which to start the Helium runtime server",
	)
	flags.StringVarP(
		&runner.heDir,
		"build-dir",
		"b",
		".helium",
		"The directory that contains the Helium build script",
	)
	return cmd

}
