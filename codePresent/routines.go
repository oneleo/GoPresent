package main

import (
	"fmt"
	"time"
)

func say(s string) { // OMIT START
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
} // OMIT END
