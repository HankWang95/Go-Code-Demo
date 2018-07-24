package book

import (
	"log"
	"time"
)

func BigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(2 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
