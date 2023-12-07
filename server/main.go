package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	p := flag.Int("port", 8080, "port")
	flag.Parse()

	port := fmt.Sprintf(":%d", *p)

	fmt.Printf("Starting server on port %s\n", port)

	s := NewServer(port)
	defer s.Stop()

	log.Fatal(s.Start())
}
