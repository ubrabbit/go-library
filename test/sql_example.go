package main

import (
	log "go-library/log"
	. "go-library/test/common"
)

const (
	USERNAME string = "umail"
	PASSWORD string = "123456"
	HOST     string = "192.168.1.78"
	PORT     int    = 6033
	DBNAME   string = "umail"
)

func main() {
	log.Info("start")

	db := ExampleConnect(USERNAME, PASSWORD, HOST, PORT, DBNAME)
	ExamplePing(db)
	ExampleTable(db)
	ExampleInsert(db)
	ExampleQuery(db)
}
