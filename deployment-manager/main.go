package main

import (
    "fmt"
    "os"
    "github.com/urfave/cli/v2"
	"deployment-manager/dm_commands"
)

func main() {
    app := &cli.App{
        Name:  "deployment-manager",
        Usage: "Manage your deployments",
        Commands: []*cli.Command{
            dm_commands.InstallCommand(),
            dm_commands.UpgradeCommand(),
            dm_commands.RollbackCommand(),
            dm_commands.UninstallCommand(),
            dm_commands.HelpCommand(),
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        fmt.Println(err)
    }
}

