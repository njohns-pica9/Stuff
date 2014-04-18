package main

import (
    "fmt"
)

type MdpaTask struct {
    db_conn *DbContainer
    Config  Config
}

func NewTask(db_conn *DbContainer, config Config) MdpaTask {
    var task = MdpaTask{db_conn, config}

    return task
}

func (t *MdpaTask) RunTask() {
    source_container := t.db_conn.GetContainerTagByName(t.Config.SourceContainer)
    dest_container := t.db_conn.GetContainerTagByName(t.Config.DestinationContainer)

    fmt.Println("Source Container: ", source_container)
    fmt.Println("Destination Container: ", dest_container)

    datapoints := t.db_conn.GetDatapointByNameAndContainerTag(t.Config.DatapointTest)

    for k, v := range datapoints {
        fmt.Printf("Key: %s Value: %s", k, v)
    }
}
