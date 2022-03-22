package ocelotLogger

import (
	"fmt"
	"log"
)

func DebugLine(key string, value interface{}) {
	log.Printf("\n**************\n %s: %+v \n**************\n", key, value)
}

func DebugSlice(key string, values ...interface{}) {
	for _, value := range values {
		log.Printf("\n**************\n %s: %+v \n**************\n", key, value)
	}
}

func ErrorLogger(errArg error) {
	log.Println(fmt.Sprintf("ERROR: %+v", errArg))
}

func InfoLogger(message string) {
	log.Printf(fmt.Sprintf("INFO: %v", message))
}
