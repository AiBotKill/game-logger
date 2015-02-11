package main

import (
	"log"
	"time"
)

var serviceId string

func main() {
	serviceId = Uuid()
	if err := startNats(); err != nil {
		log.Panicln("Can connect or start to gnatsd:", err.Error())
	}

	_, err := natsEncodedConn.Subscribe(">", func(subj string, reply string, msg []byte) {
		log.Println(string(msg))
	})
	if err != nil {
		log.Panic(err.Error())
	}

	// Logging server stats until the server is stopped.
	for {
		<-time.After(time.Second * 1)
		err := natsEncodedConn.Publish("ping", map[string]interface{}{
			"ping":      "logger",
			"serviceId": serviceId,
			"time":      time.Now().String(),
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}
