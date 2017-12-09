package actions

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/arschles/helium/pkg/config"
	log "github.com/arschles/helium/pkg/log/human"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/spf13/cobra"
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
	cl, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("couldn't connect to the Docker daemon (%s)", err)
	}
	log.Info("running job %s from %s", jobName, fileName)

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working dir")
	}
	log.Debug("working directory %s", wd)

	// TODO: check for image
	// if _, err = cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{}); err != nil {
	// 	panic(err)
	// }
	ctx := context.Background()
	resp, err := cl.ContainerCreate(ctx, &container.Config{
		Image: job.Image,
		Cmd:   []string{"sh", "-c", strings.Join(job.Tasks, " && ")},
	}, nil, nil, "") // TODO: mounts
	if err != nil {
		return fmt.Errorf("creating container (%s)", err)
	}

	if err := cl.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("starting container (%s)", err)
	}

	out, err := cl.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)

	if _, err := cl.ContainerWait(ctx, resp.ID); err != nil {
		return fmt.Errorf("failed execution (%s)", err)
	}

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
