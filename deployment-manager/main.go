package main

import (
    "deployment-manager/cmd"
    "fmt"
    "os"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
