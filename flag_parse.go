package main

import (
    "flag"
)

var Arguments struct {
    ConfigFile string
}

func init() {
    flag.StringVar(&Arguments.ConfigFile, "config", "config.json", "The configuraton file to use")
}
