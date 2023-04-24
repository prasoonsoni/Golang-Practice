package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// we have to always use the same format, this is given in documentation
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Creating our own date
	createdDate := time.Date(2002, time.March, 28, 12, 0, 0, 0, time.UTC)
	fmt.Println("My Created Date: ", createdDate.Format("01-02-2006 Monday 15:04:05"))
}
