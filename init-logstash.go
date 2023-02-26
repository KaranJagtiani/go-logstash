package logstash_logger

func Init(hostname string, port int, connection_type string, timeout int) *logstash {
	l := logstash{}
	l.host = hostname
	l.port = port
	l.connectionType = connection_type
	l.timeout = timeout
	return &l
}
