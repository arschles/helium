package actions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/kubehelium/helium/cli/config"
	"github.com/kubehelium/helium/cli/util/docker"
	log "github.com/kubehelium/helium/pkg/log/human"
)

// Build is the CLI handler for 'he build'
func Build(c *cli.Context) {
	cfg := config.ReadOrDie(c.String(FlagConfigFile))
	paths := PathsOrDie()

	dockerClient := docker.ClientOrDie()
	imgStr := fmt.Sprintf("%s:%s", cfg.Build.ImageName, cfg.Build.ImageTag)
	img, err := docker.ParseImageFromName(imgStr)
	if err != nil {
		log.Err("error parsing docker image %s (%s)", imgStr, err)
		os.Exit(1)
	}

	rmContainerCh := make(chan func())
	stdOutCh := make(chan docker.Log)
	stdErrCh := make(chan docker.Log)
	exitCodeCh := make(chan int)
	errCh := make(chan error)

	projName := filepath.Base(paths.CWD)
	binaryName := cfg.Build.GetOutputBinary(projName)
	go docker.Run(
		dockerClient,
		img,
		"build",
		paths.CWD,
		docker.ContainerGopath(cfg.Build.Gopath, paths.PackageName),
		fmt.Sprintf("go build -o %s .", binaryName),
		cfg.Build.Env,
		rmContainerCh,
		stdOutCh,
		stdErrCh,
		exitCodeCh,
		errCh,
	)

	for {
		select {
		case rmContainerFn := <-rmContainerCh:
			defer rmContainerFn()
		case l := <-stdOutCh:
			log.Info("%s", l)
		case l := <-stdErrCh:
			log.Warn("%s", l)
		case err := <-errCh:
			log.Err("%s", err)
			return
		case i := <-exitCodeCh:
			log.Info("exited with code %d", i)
			return
		}
	}

}
