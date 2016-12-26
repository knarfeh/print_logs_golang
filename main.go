package main

import (
	"fmt"
	"runtime"
	"time"
)

func loop(done chan bool) {
	for i := 0; i < 2500; i++ {
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")
		fmt.Print("balabalabalabalabalabalabalabalabalabalabalabalabala\n")

	}
	done <- true
}

func main() {
	for {
		t1 := time.Now()
		runtime.GOMAXPROCS(runtime.NumCPU())
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

		time.Sleep(5 * time.Second)
	}

}
