package main

import (
	log "go-library/log"
	. "go-library/test/example"
)

func main() {
	log.Info("start")

	ExampleIP()
	ExampleZoneID()
}
