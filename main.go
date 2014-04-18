package main

import (
    "flag"
    "fmt"
)

func main() {
    flag.Parse()
    config := NewConfig(Arguments.ConfigFile)

    fmt.Println("Hello!")

    db_container := newConnection(config.ConnString)
    defer db_container.Close()

    mdpa := NewTask(db_container, config)
    mdpa.RunTask()
}
