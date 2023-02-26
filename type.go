package logstash_logger

import "net"

type logstash struct {
	host           string
	port           int
	connectionType string
	timeout        int
	connection     net.Conn
}
