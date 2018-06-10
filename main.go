package main

import "fmt"

func main() {

	v := "i m a string dude!!"
	describe(v)

	nyobi := false
	describe(nyobi)
}

func describe(i interface{}) {
	fmt.Printf("%v, %T\n", i, i)
}
