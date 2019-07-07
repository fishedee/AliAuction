package main

import (
	"fmt"
)

//go:generate mock ^./model/.*/.*(ao|db)\.go$ ^.*(Ao|Db)$
func main() {
	fmt.Println("Hello World")
}
