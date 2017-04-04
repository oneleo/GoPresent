// +build OMIT
package main

import "fmt"

func main() {
	printHello()
}

func printHello() { // OMIT START
	fmt.Println("Hello, 世界!")
	return
} // OMIT END
