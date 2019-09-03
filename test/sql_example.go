package main

import (
	log "go-library/log"
	. "go-library/test/example"
)

func main() {
	log.Info("start")

	db := ExampleConnect(USERNAME, PASSWORD, HOST, PORT, DBNAME)
	ExamplePing(db)
	ExampleTable(db)
	ExampleInsert(db)
	ExampleQuery(db)
}
