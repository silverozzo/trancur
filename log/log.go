package log

import (
	"log"
	"os"
)

func NewInfoLog() *log.Logger {
	ll := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	return ll
}

func NewErrorLog() *log.Logger {
	ll := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	return ll
}
