package main

import (
	"fmt"
	_ "init/a"
	_ "init/b"
)

func init() {
	fmt.Println("main init")
}
func main() {
	fmt.Println("main fuction")
}
