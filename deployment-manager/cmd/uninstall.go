package cmd

import (
    "deployment-manager/pkg/operations"
    "github.com/spf13/cobra"
    "log"
)

// UninstallCmd creates and returns the 'uninstall' command
func UninstallCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "uninstall",
        Short: "Uninstall an existing deployment",
        Run: func(cmd *cobra.Command, args []string) {
            namespace, err := cmd.Flags().GetString("namespace")
            if err != nil {
                log.Fatalf("Failed to get 'namespace': %v", err)
            }

            operations.Uninstall(namespace)
        },
    }

    // Add flags
    cmd.Flags().StringP("namespace", "n", "default", "Specify the namespace for the deployment")

    return cmd
}
