// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

// !+
func handleConn(c net.Conn, shouts chan int) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup

	for input.Scan() {

		wg.Add(1)
		go echo(c, input.Text(), 2*time.Second, &wg)
		//if input.Text() != "" {
		shouts <- 1
		//}
	}
	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	tcpConn, ok := c.(*net.TCPConn)
	if !ok {
		fmt.Println("Connection is not a TCP connection")
		return
	}
	tcpConn.CloseWrite()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		shouts := make(chan int)
		go handleConn(conn, shouts)
	Timer:
		for {
			select {
			case <-time.After(10 * time.Second):
				err := conn.Close()
				if err != nil {
					return
				}
				break Timer
			case <-shouts:
			}
		}
	}
}
