package main

import (
	"os"
	"time"
)

func main() {
	dst := os.Getenv("DST-DEPLOY")
	w := &Watcher{
		Deployment: dst,
	}
	w.Run()
	time.Sleep(100 * time.Second)
}
