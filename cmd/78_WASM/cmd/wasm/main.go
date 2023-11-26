package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Set("exposedFunction", exposedFunction())

	fmt.Println("Go Web Assembly")

	<-make(chan struct{})
}
