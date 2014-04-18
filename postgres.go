package main

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

type DbContainer struct {
    Connection *sql.DB
}

func newConnection(conn_string string) *DbContainer {
    db, err := sql.Open("postgres", conn_string)

    if err != nil {
        log.Fatal(err)
    }

    var d = DbContainer{}
    d.Connection = db

    return &d
}

func (d *DbContainer) Close() {
    d.Connection.Close()
}

func (d *DbContainer) GetContainerTagByName(container_tag string) int {
    var row int

    err := d.Connection.QueryRow("SELECT container_tag_id FROM container_tags WHERE name = $1", container_tag).Scan(&row)

    switch {
    case err == sql.ErrNoRows:
        log.Printf("No container_tags with that name.")
    case err != nil:
        log.Fatal(err)
    }

    return row
}

func (d *DbContainer) GetDatapointByNameAndContainerTag(datapoint_name string) map[string]interface{} {

    var row map[string]interface{}

    err := d.Connection.QueryRow("SELECT * FROM datapoints WHERE name = $1",
        datapoint_name).Scan(&row)

    switch {
    case err == sql.ErrNoRows:
        log.Printf("No datapoints with that name.")
    case err != nil:
        log.Fatal(err)
    }

    return row
}
