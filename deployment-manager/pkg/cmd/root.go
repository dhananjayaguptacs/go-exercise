package cmd

import (
    "github.com/spf13/cobra"
)

// RootCmd represents the base command
var RootCmd = &cobra.Command{
    Use:   "deployment-manager",
    Short: "A CLI to manage your deployments",
    Long:  "Deployment Manager is a CLI tool that helps you install, upgrade, rollback, and uninstall deployments.",
}

func init() {
    RootCmd.CompletionOptions.DisableDefaultCmd = true

    // Add child commands
    RootCmd.AddCommand(InstallCmd())
    RootCmd.AddCommand(UpgradeCmd())
    RootCmd.AddCommand(RollbackCmd())
    RootCmd.AddCommand(UninstallCmd())
}
