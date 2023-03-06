package logstash_logger

import (
	"fmt"
	"log"
	"net"
	"time"
)

func (l *Logstash) initLoggerConnection() {
	logstashAddr := fmt.Sprintf("%s:%d", l.host, l.port)
	if l.connectionType == "udp" {
		netConnTemp, err := net.Dial("udp", logstashAddr)
		if err != nil {
			log.Printf("[Logstash Logger] Failed to connect to Logstash: %s\n", err)
			return
		}
		l.connection = netConnTemp
	} else if l.connectionType == "tcp" {
		netConnTemp, err := net.Dial("tcp", logstashAddr)
		if err != nil {
			log.Printf("[Logstash Logger] Failed to connect to Logstash: %s\n", err)
			return
		}
		l.connection = netConnTemp
	} else {
		log.Println("[Logstash Logger] Invalid 'ConnectionType' provided. Please provide one of the following: tcp, udp")
		return
	}

	l.setTimeout()
}

func (l *Logstash) setTimeout() {
	deadline := time.Now().Add(time.Duration(l.timeout) * time.Second)
	l.connection.SetDeadline(deadline)
	l.connection.SetWriteDeadline(deadline)
	l.connection.SetReadDeadline(deadline)
}
