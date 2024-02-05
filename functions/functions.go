package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 5
	// // funcion anonica
	// y := func() int {
	// 	return x * 2
	// }()
	// fmt.Println(y)
	c := make(chan int)
	go func() {
		fmt.Println("Starting Function")
		time.Sleep(5 * time.Second)
		fmt.Println("End Function")
		c <- 1
	}()
	<-c
}
