package main

import (
	"fmt"
	"time"
)

func main() {

	//struct{} を使用する場合は空の構造体なので close が利用できる
	quit := make(chan struct{})

	go func(){
		fmt.Println("goroutine")
		time.Sleep(2 * time.Second)
		close(quit)
		fmt.Println("close")
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)	
		select {
			case <-quit:
				fmt.Println("done")
			default:
				fmt.Println("default")
		}
	} 

	

}