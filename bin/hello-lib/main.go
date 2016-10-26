package main

import (
	"fmt"
	"github.com/cwchentw/hello-lib"
	"os"
)

func main() {
	// Arguments without the program name
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println(hello.Hello("World"))
	} else {
		fmt.Println(hello.Hello(args[0]))
	}
}
