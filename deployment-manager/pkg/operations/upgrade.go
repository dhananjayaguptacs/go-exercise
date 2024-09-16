package operations

import "fmt"

// Upgrade handles upgrading a deployment
func Upgrade(version string) {
    fmt.Printf("Upgrading deployment to version '%s'\n", version)
}
