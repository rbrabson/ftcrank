package main

import (
	"os"

	"github.com/rbrabson/ftcrank/ftcdata"
	"github.com/urfave/cli/v2"
)

var (
	appName  = "update"
	Version  = "0.1.0"
	Revision = "v1beta1"
	usage    = "FTC match prediction utility"
)

var (
	// flags are the set of flags supported by the prediction application
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			EnvVars:  []string{"FTC_USERNAME"},
			Usage:    "Username used to log into the FTC API website",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "password",
			Aliases:  []string{"p"},
			EnvVars:  []string{"FTC_AUTHORIZATION_KEY"},
			Usage:    "The authorization token used as the password when logging into the FTC API website",
			Required: true,
		},
		&cli.StringFlag{
			Name:        "server",
			Aliases:     []string{"s"},
			EnvVars:     []string{"FTC_SERVER"},
			Usage:       "API server URL",
			DefaultText: "https://ftc-api.firstinspires.org/v2.0",
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "dir",
			Aliases:     []string{"d"},
			EnvVars:     []string{"STORAGE_DIRECTORY"},
			Usage:       "Base directory under which data is stored",
			DefaultText: "",
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "season",
			Aliases:     []string{"y"},
			EnvVars:     []string{"FTC_SEASON"},
			Usage:       "Season for the FTC data",
			DefaultText: "2023",
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "event",
			Aliases:     []string{"e"},
			Usage:       "FTC event from which to predict the results",
			DefaultText: "x",
		},
		&cli.StringFlag{
			Name:  "log",
			Usage: "Log level to be used when logging messages",
			Value: "Warn",
		},
	}
)

func runApp(cli *cli.Context) error {
	ftcdata.LoadAll()
	eventCode := cli.String("event")
	if eventCode == "" {
		ftcdata.RetrieveAll()
	} else {
		ftcdata.UpdateAll(eventCode)
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:    appName,
		Version: Version + "+" + Revision,
		Flags:   flags,
		Usage:   usage,
		Action:  runApp,
	}

	app.Run(os.Args)
}
