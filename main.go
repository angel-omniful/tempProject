package main

import (
    "github.com/omniful/go_commons/log"
    "fmt"
)

func main() {
    fmt.Println("this is the change for second commit")
    log.Infof("OMS service is starting up...")

    // Example logs at various levels
    log.Debugf("This is a debug message: %v", "debug-info")
    log.Warnf("This is a warning!")
    log.Errorf("This is an error log with code %d", 500)

    // Normally you'd start server / worker here
    // server.Start()
}