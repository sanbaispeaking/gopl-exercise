package main

import "fmt"

func noReturnStatement() (out int) {
	// Deferred function get executed after caller paniced
	defer func () {
		if p := recover(); p != nil {
			out = p.(int)
		}
	}()

	panic(42)
}


func main()  {
	fmt.Printf("The answer is %v\n", noReturnStatement())
}