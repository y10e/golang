package main

import (
	"log"
	"time"
)

func f() {
    log.Println("goroutine")
}

func main() {
    
	//goroutine
	go f()
	
	//goroutine(無名関数)
	go func() {
		log.Println("goroutine2")
	}()
	
    log.Println("main")
    time.Sleep(time.Second)
    log.Println("finish")
}