package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	lockFilename := flag.String("lock", "my.lock", "lock filename")
	sleepDuration := flag.Duration("sleep", time.Second, "sleep duration")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	pid := os.Getpid()
	log.Printf("program started, pid=%d", pid)
	file, err := os.OpenFile(*lockFilename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = os.Remove(*lockFilename)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("start sleeping, pid=%d", pid)
	time.Sleep(*sleepDuration)
	log.Printf("finish sleeping, pid=%d", pid)
}
