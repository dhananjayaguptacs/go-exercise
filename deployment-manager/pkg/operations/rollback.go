package operations

import "fmt"

// Rollback handles rolling back a deployment
func Rollback(version string) {
    fmt.Printf("Rolling back deployment to version '%s'\n", version)
}
