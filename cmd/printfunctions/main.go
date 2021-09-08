package main

import (
	"fmt"

	"github.com/webmachinedev/src/functions"
)

func main() {
	f, err := functions.AllFunctions()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(f))
}
