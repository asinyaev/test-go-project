package main

import (
	"fmt"
	"github.com/asinyaev/mymath"
	mymath2 "github.com/asinyaev/mymath/v2"
)

func main() {
	fmt.Println(mymath.Add(2, 2))
	fmt.Println(mymath.Sub(2, 2))
	fmt.Println(mymath2.Add(1, 2, 3))
}
