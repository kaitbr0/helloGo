package main

//declare main package
import (
	"fmt"

	"rsc.io/quote"
) //import fmt package, functions for formatting text

func main() {
    fmt.Println(quote.Opt())
}

//go mod tidy to import the package
// go run .
