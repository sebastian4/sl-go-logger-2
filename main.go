package slgologger2

import "log"

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func LogWarn(message string) {
	log.Printf("WARN2 - %v", message)
}

func LogError(message string) {
	log.Printf("ERROR - %v", message)
}
