package cmd

import (
    "deployment-manager/pkg/operations"
    "github.com/spf13/cobra"
    "log"
)

// UpgradeCmd creates and returns the 'upgrade' command
func UpgradeCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "upgrade",
        Short: "Upgrade an existing deployment",
        Run: func(cmd *cobra.Command, args []string) {
            version, err := cmd.Flags().GetString("version")
            if err != nil {
                log.Fatalf("Failed to get 'version': %v", err)
            }

            operations.Upgrade(version)
        },
    }

    // Add flags
    cmd.Flags().StringP("version", "v", "latest", "Specify the version to upgrade to")

    return cmd
}
