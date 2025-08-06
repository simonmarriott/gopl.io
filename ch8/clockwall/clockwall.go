// clockwall listens to multiple clock servers concurrently.
package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type clock struct {
	name, host string
	port       int
}

func (c *clock) watch(w io.Writer, r net.Conn) {
	b := make([]byte, 1024)
	_, _ = r.Read(b)
	_, _ = fmt.Fprintf(w, "%s %s", c.name, b)
	_ = r.Close()
	return
}

func main() {
	if len(os.Args) == 1 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: clockwall NAME=HOST:PORT ...")
		os.Exit(1)
	}
	clocks := make([]*clock, 0)
	for _, a := range os.Args[1:] {
		fields := strings.Split(a, "=")
		address := strings.Split(fields[1], ":")

		if len(fields) != 2 || len(address) != 2 {
			_, _ = fmt.Fprintf(os.Stderr, "bad arg: %s\n", a)
			os.Exit(1)
		}
		port, _ := strconv.Atoi(address[1])
		clocks = append(clocks, &clock{fields[0], address[0], port})
	}
	for {
		for n, c := range clocks {
			conn, err := net.Dial("tcp", c.host+":"+strconv.Itoa(c.port))
			if err != nil {
				continue
			}
			for _ = range n {
				fmt.Printf("\t\t")
			}
			c.watch(os.Stdout, conn)

		}
		fmt.Printf("\r")
		time.Sleep(time.Second)
	}

}
