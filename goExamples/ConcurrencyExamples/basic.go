package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("start")
	go Process()
	time.Sleep(time.Millisecond*10)
	fmt.Println("Done")
}

func Process(){
	fmt.Println("Processing")
}
