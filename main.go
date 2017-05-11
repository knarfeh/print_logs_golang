package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

func loop(done chan bool) {
	printCount := getEnvInt("PRINT_COUNT", 1)
	printContent := getEnv("PRINT_CONTENT", "balabalabalabalabalabalabalabalabalabalabalabalabala")
	for i := 0; i < printCount; i++ {
		fmt.Println(printContent)
		fmt.Println(printContent)
		fmt.Println(printContent)
		fmt.Println(printContent)
		fmt.Println(printContent)
		fmt.Println(printContent)
		fmt.Println(printContent)
		fmt.Println(printContent)
		fmt.Println(printContent)
		fmt.Println(printContent)
	}
	done <- true
}

// getEnv get an environment key as string, returns the default if not found
func getEnv(key string, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultVal
	}
	return value
}

// getEnvInt get an env key as int, returns the default if note found
func getEnvInt(key string, defaultVal int) int {
	valueStr := getEnv(key, strconv.Itoa(defaultVal))
	res, err := strconv.Atoi(valueStr)
	if err != nil {
		res = defaultVal
	}
	return res
}

func main() {
	for {
		t1 := time.Now()
		runtime.GOMAXPROCS(runtime.NumCPU())
		done := make(chan bool)
		go loop(done)
		go loop(done)

		<-done
		<-done
		t2 := time.Now()
		fmt.Println("time span: ", t2.Sub(t1))

		time.Sleep(10 * time.Second)
	}
}
