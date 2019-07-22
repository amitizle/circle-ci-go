package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("hello, world, run time: %s\n", t.Format(time.RFC3339))
}
