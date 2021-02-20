package main

import (
	"fmt"
	"log"
)

func printa(a chan string) {
    a <- "a"
}

func printb(b chan string) {
    b <- "b"
}

func printc(c chan string) {
    c <- "c"
}

func main() {
    a := make(chan string)
    b := make(chan string)
    c := make(chan string)
    log.Print("started")
    go printa(a)
    go printb(b)
    go printc(c)
    for i := 0; i < 3; i++ {
        select {
        case msga := <-a:
            fmt.Println("finished", msga)
        case msgb := <-b:
            fmt.Println("finished", msgb)
        case msgc := <-c:
            fmt.Println("finished", msgc)

        }
    }

    log.Print("finished")

}