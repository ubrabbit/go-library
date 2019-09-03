package example

import (
	log "go-library/log"
	. "go-library/net/ip"
)

func ExampleIP() {

	log.Info("InternalIP: %s", InternalIP())
	// InetAtoN
	log.Info("InetAtoN(183.131.11.57): %d", InetAtoN("183.131.11.57"))
	log.Info("InetAtoN(183.131.11): %d", InetAtoN("183.131.11"))
	log.Info("InetAtoN(0:0:0:0:0:0:0:1): %d", InetAtoN("0:0:0:0:0:0:0:1"))
	// InetNtoA
	log.Info("InetNtoA(84549632): %s", InetNtoA(84549632))

	l2 := ExternalIP()
	log.Info("ExternalIP: %v", l2)
}

func ExampleZoneID() {
	// ZoneID
	log.Info("ZoneID: %d", ZoneID("中国", "福建", "莆田"))
}
