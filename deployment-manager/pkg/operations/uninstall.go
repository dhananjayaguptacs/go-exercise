package operations

import "fmt"

// Uninstall handles the uninstallation of a deployment
func Uninstall(namespace string) {
    fmt.Printf("Uninstalling deployment in namespace '%s'\n", namespace)
}
