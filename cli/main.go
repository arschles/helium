package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/kubehelium/helium/cli/actions"
	log "github.com/kubehelium/helium/pkg/log/human"
)

func main() {
	app := cli.NewApp()
	app.Name = "helium"
	app.Usage = "helium is your one-stop shop for building and deploying Kubernetes apps"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  actions.FlagDebug,
			Usage: "Enable verbose debugging output",
		},
		cli.StringFlag{
			Name:  actions.FlagConfigDir,
			Usage: "The location of the config directory",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "build",
			Aliases:     []string{"b"},
			Usage:       "Build your project",
			Description: "This command runs the equivalent of 'go build -o $CURRENT_DIR_NAME'",
			Action:      actions.Build,
		},
		// {
		// 	Name:        "docker",
		// 	Aliases:     []string{"d"},
		// 	Description: "Docker related actions",
		// 	Subcommands: []cli.Command{
		// 		{
		// 			Name:        "build",
		// 			Aliases:     []string{"db"},
		// 			Usage:       "Build a Docker image for your project",
		// 			Description: "This command runs the equivalent of 'docker build -t $IMG_NAME $DOCKERFILE_DIR'",
		// 			Action:      actions.DockerBuild,
		// 		},
		// 		{
		// 			Name:        "doc-push",
		// 			Aliases:     []string{"dp"},
		// 			Usage:       "Push the Docker image for your project",
		// 			Description: "This command runs the equivalent of 'docker push $IMG_NAME'",
		// 			Action:      actions.DockerPush,
		// 		},
		// 	},
		// },
		// {
		// 	Name:    "custom",
		// 	Aliases: []string{"c"},
		// 	Flags: []cli.Flag{
		// 		cli.BoolFlag{
		// 			Name: actions.ListCustomFlag,
		// 		},
		// 	},
		// 	Usage:       "Run a custom target",
		// 	Description: "Run a custom build target.",
		// 	Action:      actions.Custom,
		// },
	}

	app.Before = func(c *cli.Context) error {
		log.IsDebugging = c.Bool(actions.FlagDebug)
		return nil
	}

	app.Run(os.Args)
}
