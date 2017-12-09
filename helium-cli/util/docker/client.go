package docker

import (
	"os"

	docker "github.com/fsouza/go-dockerclient"
	log "github.com/arschles/helium/pkg/log/human"
)

// ClientOrDie creates a new Docker client. If one couldn't be created, logs and error and exits with status code 1
func ClientOrDie() *docker.Client {
	cl, err := docker.NewClientFromEnv()
	if err != nil {
		log.Err("creating new docker client (%s)", err)
		os.Exit(1)
	}
	return cl
}
