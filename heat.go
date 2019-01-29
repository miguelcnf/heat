package main

import (
	"crypto/sha512"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	cores := runtime.NumCPU()

	signals := make(chan os.Signal, 1)
	exit := make(chan bool)

	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	for i := 0; i < cores; i++ {
		go getHot(signals, exit)
	}

	<-exit
}

func getHot(signals chan os.Signal, exit chan bool) {
	hasher := sha512.New()
	random := string(rand.Int63())
	for {
		select {
		case <-signals:
			exit <-true
		default:
			_, _ = hasher.Write([]byte(random))
		}
	}
}
