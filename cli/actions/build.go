package actions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	// "path/filepath"

	"github.com/codegangsta/cli"
	// "github.com/kubehelium/helium/cli/util/docker"
	log "github.com/kubehelium/helium/pkg/log/human"
	"github.com/kubehelium/helium/pkg/runtime"
)

// Build is the CLI handler for 'he build'
func Build(c *cli.Context) {
	// event := c.String("target")
	// dockerClient := docker.ClientOrDie()
	log.Info("build")
	srv := runtime.NewServer(8080)
	srvErrCh := make(chan error)
	go func() {
		log.Debug("Running the event handler server")
		srvErrCh <- srv.Run()
	}()

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
		log.Err("error getting current working dir")
		os.Exit(1)
	}
	heliumDir := fmt.Sprintf("%s/.helium", wd)
	log.Info("using script dir %s", heliumDir)
	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-v",
		fmt.Sprintf("%s:/helium", heliumDir),
		"-e",
		"HELIUM_SCRIPTS_DIR=/helium",
		"-e",
		"HELIUM_TARGET_DIR=/home/src/dist/src",
		imgStr,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Info(strings.Join(cmd.Args, " "))
	if err := cmd.Run(); err != nil {
		log.Err("error running command (%s)", err)
		os.Exit(1)
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

}
