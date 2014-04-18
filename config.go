package main

import (
    "encoding/json"
    "os"
)

type Config struct {
    ConnString           string `json:"db_connection"`
    SourceContainer      string `json:"source_container"`
    DestinationContainer string `json:"destination_container"`
    DatapointTest        string `json:"dp_test"`
}

func NewConfig(filename string) Config {
    c := Config{}
    c.loadConfig(filename)

    return c
}

func (c *Config) loadConfig(filename string) error {

    configFile, err := os.Open(filename)

    if err != nil {
        return err
    }

    defer configFile.Close()

    err = json.NewDecoder(configFile).Decode(c)

    if err != nil {
        return err
    }

    return nil
}
