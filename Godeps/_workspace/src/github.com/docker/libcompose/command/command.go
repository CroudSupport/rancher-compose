package command

import (
	"github.com/codegangsta/cli"
	"github.com/docker/libcompose/app"
	"github.com/docker/libcompose/project"
)

func CreateCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:   "create",
		Usage:  "Create all services but do not start",
		Action: app.WithProject(factory, app.ProjectCreate),
	}
}

func UpCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:   "up",
		Usage:  "Bring all services up",
		Action: app.WithProject(factory, app.ProjectUp),
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "d",
				Usage: "Do not block and log",
			},
		},
	}
}

func StartCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:   "start",
		Usage:  "Start services",
		Action: app.WithProject(factory, app.ProjectUp),
	}
}

func LogsCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:   "logs",
		Usage:  "Get service logs",
		Action: app.WithProject(factory, app.ProjectLog),
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "lines",
				Usage: "number of lines to tail",
				Value: 100,
			},
		},
	}
}

func RestartCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:   "restart",
		Usage:  "Restart services",
		Action: app.WithProject(factory, app.ProjectRestart),
	}
}

func StopCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:      "stop",
		ShortName: "down",
		Usage:     "Stop services",
		Action:    app.WithProject(factory, app.ProjectDown),
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "timeout",
				Usage: "Specify a shutdown timeout in seconds.",
				Value: 10,
			},
		},
	}
}

func ScaleCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:   "scale",
		Usage:  "Scale services",
		Action: app.WithProject(factory, app.ProjectScale),
	}
}

func RmCommand(factory app.ProjectFactory) cli.Command {
	return cli.Command{
		Name:   "rm",
		Usage:  "Delete services",
		Action: app.WithProject(factory, app.ProjectDelete),
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "force,f",
				Usage: "Allow deletion of all services",
			},
		},
	}
}

func CommonFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name: "verbose,debug",
		},
		cli.StringFlag{
			Name:   "file,f",
			Usage:  "Specify an alternate compose file (default: docker-compose.yml)",
			Value:  "docker-compose.yml",
			EnvVar: "COMPOSE_FILE",
		},
		cli.StringFlag{
			Name:  "project-name,p",
			Usage: "Specify an alternate project name (default: directory name)",
		},
	}
}

func Populate(context *project.Context, c *cli.Context) {
	context.ComposeFile = c.GlobalString("file")
	context.ProjectName = c.GlobalString("project-name")

	if c.Command.Name == "logs" {
		context.Log = true
	} else if c.Command.Name == "up" {
		context.Log = !c.Bool("d")
	} else if c.Command.Name == "stop" {
		context.Timeout = c.Int("timeout")
	}
}
