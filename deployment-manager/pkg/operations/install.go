package operations

import "fmt"

// Install handles the installation logic for a deployment
func Install(namespace string, timeout int, ignore bool) {
    fmt.Printf("Installing deployment in namespace '%s' with timeout %d seconds and ignore: %t\n", namespace, timeout, ignore)
}
