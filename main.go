// main.go

package main

import (
	"fmt"
	"myproject/mypackage"
)

func main() {
	sum := mypackage.Add(2, 3)
	fmt.Println("Sum:", sum)
}
