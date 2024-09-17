package cmd

import (
    "deployment-manager/pkg/operations"
    "github.com/spf13/cobra"
    "log"
)

// RollbackCmd creates and returns the 'rollback' command
func RollbackCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "rollback",
        Short: "Rollback to a previous version of the deployment",
        Run: func(cmd *cobra.Command, args []string) {
            version, err := cmd.Flags().GetString("version")
            if err != nil {
                log.Fatalf("Failed to get 'version': %v", err)
            }

            operations.Rollback(version)
        },
    }

    // Add flags
    cmd.Flags().StringP("version", "v", "", "Specify the version to rollback to")

    return cmd
}
