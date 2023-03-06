package logstash_logger

func (l *Logstash) Log(payload map[string]interface{}) {
	l.pushJsonMessage(payload)
}

func (l *Logstash) Info(payload map[string]interface{}) {
	payload["severity"] = "INFO"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *Logstash) Debug(payload map[string]interface{}) {
	payload["severity"] = "DEBUG"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *Logstash) Warn(payload map[string]interface{}) {
	payload["severity"] = "WARN"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *Logstash) Error(payload map[string]interface{}) {
	payload["severity"] = "ERROR"
	l.pushJsonMessage(payload)
	delete(payload, "severity")
}

func (l *Logstash) LogString(stringMsg string) {
	l.pushStringMessage(stringMsg)
}
