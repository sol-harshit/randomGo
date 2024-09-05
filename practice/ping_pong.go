package main

import (
	"fmt"
	"time"
)

var ping chan int = make(chan int)
var pong chan int = make(chan int)

func main() {

	go pingNow(ping, pong)
	ping <- 1
	select {}
}

func pingNow(ping chan int, pong chan int) {
	<-ping
	fmt.Println("Ping")
	time.Sleep(time.Millisecond * 200)
	go pongNow(ping, pong)
	pong <- 1
}

func pongNow(ping chan int, pong chan int) {
	<-pong
	fmt.Println("pong")
	time.Sleep(time.Millisecond * 200)
	go pingNow(ping, pong)
	ping <- 1
}
