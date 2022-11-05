package main

import (
	"flag"
	"fmt"
)

var (
	port = flag.Int("port", 5300, "server port")
)

func main() {
	flag.Parse()

	fmt.Printf("server port is: %d\n", *port)
}
