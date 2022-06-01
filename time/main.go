package main

import (
	"fmt"
	"time"
)
//define a channel
var chs chan int

// time.After memory leak
func Get() {
	for {
		select {
			case v := <- chs:
				fmt.Printf("print:%v\n", v)
			case <- time.After(3 * time.Minute):
				fmt.Printf("time.After:%v", time.Now().Unix())
		}
	}
}

// fixed version
func Get2() {
	delay := time.NewTimer(3 * time.Minute)

	defer delay.Stop()

	for {
		delay.Reset(3 * time.Minute)

		select {
			case v := <- chs:
				fmt.Printf("print:%v\n", v)
			case <- delay.C:
				fmt.Printf("time.After:%v", time.Now().Unix())
		}
	}
}

func Put() {
	var i = 0
	for {
		i++
		chs <- i
	}
}

func main() {
	chs = make(chan int, 100)
	go Put()
	Get2()
}