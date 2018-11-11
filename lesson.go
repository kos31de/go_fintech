package main

import (
	"fmt"
	"time"
)

func init () {
	fmt.Println("Init", time.Now())
}

func bazz () {
	fmt.Println("Bazz")
}

func main () {
	bazz()
	fmt.Println("Hello00000")
}
