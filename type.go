package logstash_logger

import "net"

type Logstash struct {
	host           string
	port           int
	connectionType string
	timeout        int
	connection     net.Conn
}
