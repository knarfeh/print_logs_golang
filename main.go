package main

import (
	"fmt"
	"runtime"
	"time"
)

func loop(done chan bool) {
	for i := 0; i < 25000; i++ {
		fmt.Printf("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
	}
	done <- true
}

func main() {
	t1 := time.Now()
	runtime.GOMAXPROCS(4)
	done := make(chan bool)
	go loop(done)
	go loop(done)
	go loop(done)
	go loop(done)

	<-done
	<-done
	<-done
	<-done
	t2 := time.Now()
	fmt.Println("time span: ", t2.Sub(t1))
}
