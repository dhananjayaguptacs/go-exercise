package dm_commands

import (
    "github.com/urfave/cli/v2"
)

// InstallCommand returns the command configuration for 'install'
func InstallCommand() *cli.Command {
    return &cli.Command{
        Name:    "install",
        Aliases: []string{"ii"},
        Usage:   "Install a new deployment",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "namespace",
                Value:   "default",
                Usage:   "Specify the namespace for the deployment",
            },
            &cli.IntFlag{
                Name:    "timeout",
                Value:   30,
                Usage:   "Specify the timeout for the deployment in seconds",
            },
            &cli.BoolFlag{
                Name: "ignore",
                Required: true,
                Usage:   "[Required] Specify the ignore value in bool",
            },
        },
        Action: func(c *cli.Context) error {
            return install(c.String("namespace"), c.Int("timeout"), c.Bool("ignore"))
        },
    }
}

// UpgradeCommand returns the command configuration for 'upgrade'
func UpgradeCommand() *cli.Command {
    return &cli.Command{
        Name:    "upgrade",
        Aliases: []string{"up"},
        Usage:   "Upgrade an existing deployment",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "version",
                Value:   "latest",
                Usage:   "Specify the version to upgrade to",
            },
        },
        Action: func(c *cli.Context) error {
            return upgrade(c.String("version"))
        },
    }
}

// RollbackCommand returns the command configuration for 'rollback'
func RollbackCommand() *cli.Command {
    return &cli.Command{
        Name:    "rollback",
        Aliases: []string{"rb"},
        Usage:   "Rollback to a previous deployment",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "version",
                Usage:   "Specify the version to rollback to",
            },
        },
        Action: func(c *cli.Context) error {
            return rollback(c.String("version"))
        },
    }
}

// UninstallCommand returns the command configuration for 'uninstall'
func UninstallCommand() *cli.Command {
    return &cli.Command{
        Name:    "uninstall",
        Aliases: []string{"un"},
        Usage:   "Uninstall a deployment",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "namespace",
                Value:   "default",
                Usage:   "Specify the namespace for the deployment",
            },
        },
        Action: func(c *cli.Context) error {
            return uninstall(c.String("namespace"))
        },
    }
}

// HelpCommand returns the command configuration for 'help'
func HelpCommand() *cli.Command {
    return &cli.Command{
        Name:    "help",
        Aliases: []string{"h"},
        Usage:   "Show help information",
        Action: func(c *cli.Context) error {
            return cli.ShowAppHelp(c)
        },
    }
}