package cmd

import (
    "deployment-manager/pkg/operations"
    "github.com/spf13/cobra"
    "log"
)

// InstallCmd creates and returns the 'install' command
func InstallCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "install",
        Short: "Install a new deployment",
        Run: func(cmd *cobra.Command, args []string) {
            namespace, err := cmd.Flags().GetString("namespace")
            if err != nil {
                log.Fatalf("Failed to get 'namespace': %v", err)
            }

            timeout, err := cmd.Flags().GetInt("timeout")
            if err != nil {
                log.Fatalf("Failed to get 'timeout': %v", err)
            }

            ignore, err := cmd.Flags().GetBool("ignore")
            if err != nil {
                log.Fatalf("Failed to get 'ignore': %v", err)
            }

            operations.Install(namespace, timeout, ignore)
        },
    }

    // Add flags
    cmd.Flags().StringP("namespace", "n", "", "Specify the namespace for the deployment (required)")
    cmd.MarkFlagRequired("namespace")
    cmd.Flags().IntP("timeout", "t", 0, "Specify the timeout for the deployment in seconds (required)")
    cmd.MarkFlagRequired("timeout")
    cmd.Flags().Bool("ignore", false, "Ignore pre-install checks (use --ignore=true or --ignore=false)")

    return cmd
}
