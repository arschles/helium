package actions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	// "github.com/arschles/helium/cli/util/docker"
	"github.com/arschles/helium/pkg/config"
	log "github.com/arschles/helium/pkg/log/human"
)

type runner struct {
	port  int
	heDir string
}

func (r *runner) run(c *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Usage: he run <job>")
	}
	jobName := args[0]

	fileName := filepath.Join(r.heDir, "config.toml")
	_, dirErr := os.Stat(fileName)
	if os.IsNotExist(dirErr) {
		return fmt.Errorf("%s does not exist", fileName)
	}
	_, cfg, err := config.Parse(fileName)
	if err != nil {
		return fmt.Errorf("couldn't parse config file (%s)", err)
	}
	job, ok := cfg.Jobs[jobName]
	if !ok {
		return fmt.Errorf("Job %s doesn't exist in the config file")
	}
	// event := c.String("target")
	// dockerClient := docker.ClientOrDie()
	// TODO: make configurable
	log.Info("running job %s from %s", jobName, fileName)

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
	log.Debug("working directory %s", wd)

	// TODO: use the docker API directly
	// https://docs.docker.com/develop/sdk/#api-version-matrix
	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-v",
		fmt.Sprintf("%s:/wd", wd),
		job.Image,
		fmt.Sprintf(`sh -c '%s'`, strings.Join(job.Tasks, " && ")),
	)
	log.Debug(strings.Join(cmd.Args, " "))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s", err)

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
		Short: "Execute an individual job",
		RunE:  runner.run,
	}
	flags := cmd.Flags()
	flags.StringVarP(
		&runner.heDir,
		"helium-dir",
		"d",
		".helium",
		"The directory that contains the Helium config.toml",
	)
	return cmd

}
