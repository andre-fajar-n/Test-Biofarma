package main

import (
	"biofarma/cmd/server"
	"biofarma/runtime"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	appName  = "Test Biofarma"
	appUsage = "CLI to run this apps"

	flagNamePath = "path"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appUsage
	app.Commands = []*cli.Command{
		api, migration,
	}
	app.CommandNotFound = commanNotFound

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func commanNotFound(ctx *cli.Context, s string) {
	fmt.Println("Command not found. Please use command \"help\" or \"h\" to show help")
}

var (
	api = &cli.Command{
		Name:  "api",
		Usage: "Run API Server",
		Action: func(ctx *cli.Context) error {
			server.Main()
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "port",
				Usage:       "server port",
				DefaultText: "random port",
			},
			&cli.StringFlag{
				Name:        "host",
				Usage:       "server host",
				DefaultText: "localhost",
			},
		},
	}

	migration = &cli.Command{
		Name:  "migration",
		Usage: "Run migration needs",
		Subcommands: []*cli.Command{
			migrateUp, migrateDown, migrateCreateFile, migrateForce,
		},
	}

	migrateUp = &cli.Command{
		Name:  "up",
		Usage: "Migrate up migrations",
		Action: func(ctx *cli.Context) error {
			rt := runtime.NewRuntime()

			rt.MigrateUp(ctx.String(flagNamePath))

			return nil
		},
		Flags: []cli.Flag{
			&migrationFilePath,
		},
	}

	migrateDown = &cli.Command{
		Name:  "down",
		Usage: "Migrate down migrations",
		Action: func(ctx *cli.Context) error {
			rt := runtime.NewRuntime()

			rt.MigrateDown(ctx.String(flagNamePath))

			return nil
		},
		Flags: []cli.Flag{
			&migrationFilePath,
		},
	}

	migrateForce = &cli.Command{
		Name:  "force",
		Usage: "Fix latest dirty migration",
		Action: func(ctx *cli.Context) error {
			rt := runtime.NewRuntime()

			rt.ForceLastestVersion(ctx.String(flagNamePath))

			return nil
		},
		Flags: []cli.Flag{
			&migrationFilePath,
		},
	}

	migrateCreateFile = &cli.Command{
		Name:  "create",
		Usage: "Create new file migrations",
		Action: func(ctx *cli.Context) error {
			rt := runtime.NewRuntime()

			args := ctx.Args()

			if args.Len() < 1 {
				return fmt.Errorf("please insert filename")
			}

			filename := args.First()
			rt.CreateFileMigration(ctx.String(flagNamePath), filename)

			return nil
		},
		Flags: []cli.Flag{
			&migrationFilePath,
		},
	}

	migrationFilePath = cli.StringFlag{
		Name:        flagNamePath,
		Usage:       "path where migration file placed",
		DefaultText: "./internal/migrations",
		Value:       "./internal/migrations",
	}
)
