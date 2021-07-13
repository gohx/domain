package main

import (
	"fmt"

	"github.com/gohx/domain"
)

func main() {
	d := domain.NewDomain("example.com")
	fmt.Printf("%#+v\n", d)
}
