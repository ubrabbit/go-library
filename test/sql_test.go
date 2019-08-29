package main

/*
go test -cover -v -bench=.
*/

import (
	"testing"
)

import (
	log "go-library/log"
	. "go-library/test/common"
)

func TestMySQL(t *testing.T) {
	log.Info("start")

	db := ExampleConnect(USERNAME, PASSWORD, HOST, PORT, DBNAME)
	ExamplePing(db)
	ExampleTable(db)
	ExampleInsert(db)
	ExampleQuery(db)
}
