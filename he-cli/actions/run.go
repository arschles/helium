package actions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/arschles/helium/pkg/config"
	log "github.com/arschles/helium/pkg/log/human"
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
		return fmt.Errorf("Job %s doesn't exist in the config file", jobName)
	}

	log.Info("running job %s from %s", jobName, fileName)

	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working dir")
	}
	log.Debug("working directory %s", wd)

	for i, task := range job.Tasks {
		log.Debug(task)
		spl := strings.Split(task, " ")
		if len(spl) < 1 {
			continue
		}
		first := spl[0]
		cmd := exec.Command(first, spl[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Err("task %d failed (%s)", i, err)
		}
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
