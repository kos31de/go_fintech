package main

import (
	"fmt"
	"os/user"
	"time"
)

func init () {
	fmt.Println("Init", time.Now())
	fmt.Println(user.Current())
}
