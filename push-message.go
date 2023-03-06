package logstash_logger

import (
	"encoding/json"
	"log"
)

func (l *Logstash) pushJsonMessage(jsonMsg map[string]interface{}) {
	l.initLoggerConnection()

	if l.connection == nil {
		log.Println("[Logstash Logger] Logger was not initialized properly, could not push message.")
		return
	}

	jsonData, err := json.Marshal(jsonMsg)

	if err != nil {
		log.Printf("[Logstash Logger] Failed to parse JSON object: %s\n", err)
	}

	_, connWriteError := l.connection.Write(jsonData)
	if connWriteError != nil {
		log.Printf("[Logstash Logger] Failed to send log JSON message: %s\n", connWriteError)
	}

	l.connection.Close()
}

func (l *Logstash) pushStringMessage(stringMsg string) {
	l.initLoggerConnection()

	if l.connection == nil {
		log.Println("[Logstash Logger] Logger was not initialized properly, could not push message.")
		return
	}

	_, connWriteError := l.connection.Write([]byte(stringMsg))
	if connWriteError != nil {
		log.Printf("[Logstash Logger] Failed to send log String message: %s\n", connWriteError)
	}

	l.connection.Close()
}
