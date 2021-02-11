package main

import (
	"fmt"
	"log"
	"time"
)

func f(c chan string) {
    time.Sleep(10 * time.Second)
    c <- "well come"
}

func main() {
    c := make(chan string)
    log.Print("started")
    go f(c)
    w1 := <-c//go f(c)実行後、即座に呼ばれていない
    log.Print("finished")
    fmt.Println(w1)
}