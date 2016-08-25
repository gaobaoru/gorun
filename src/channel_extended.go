package main

import(
    "fmt"
    "time"
)

type Sender chan<- int
type Receiver <-chan int

func main(){
    var myChannel = make(chan int, 0)
    var number = 6
    go func(){
	var sender Sender = myChannel
	sender <- number
    }()
    go func(){
	var receiver Receiver = myChannel
	fmt.Println("Received!", <-receiver)
	fmt.Println("Sent!")
    }()

    time.Sleep(time.Second)
}
