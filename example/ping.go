package main

import (
	"fmt"
	"shkabatur/naiveping"
	"time"
)

func main() {
	addr := "google.com"
	timeout := 2 * time.Second

	fmt.Println(naiveping.Ping(addr, timeout))
}
