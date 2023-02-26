package logstash_logger

func (l *logstash) Log(payload map[string]interface{}) {
	l.pushJsonMessage(payload)
}

func (l *logstash) Info(payload map[string]interface{}) {
	payload["severity"] = "INFO"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *logstash) Debug(payload map[string]interface{}) {
	payload["severity"] = "DEBUG"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *logstash) Warn(payload map[string]interface{}) {
	payload["severity"] = "WARN"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *logstash) Error(payload map[string]interface{}) {
	payload["severity"] = "ERROR"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *logstash) LogString(stringMsg string) {
	l.pushStringMessage(stringMsg)
}
