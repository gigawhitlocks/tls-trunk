package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
)

func chanFromConn(conn net.Conn) chan []byte {
	c := make(chan []byte)
	go func() {
		for {
			b := make([]byte, 1024)
			n, err := conn.Read(b)
			if n > 0 {
				res := make([]byte, n)
				copy(res, b[:n])
				c <- res
			}
			if err != nil {
				c <- nil
				break
			}
		}
	}()
	return c
}

func main() {
	certificate, err := tls.LoadX509KeyPair("/home/ian/ian.crt.pem",
		"/home/ian/ian.key.pem")
	if err != nil {
		fmt.Printf("Failed to load certs: %s\n", err)
		os.Exit(1)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
	}

	ln, err := net.Listen("tcp", ":8080")
	fmt.Printf("Listening on :8080\n")
	if err != nil {
		fmt.Printf("Failed to listen: %s\n", err)
		os.Exit(1)
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Printf("Failed to accept: %s\n", err)
			continue
		}

		go func(c net.Conn) {
			client, err := tls.Dial("tcp", "secret.theknown.net:443", config)
			if err != nil {
				fmt.Printf("Failed to connect: %s\n", err)
				return
			}

			upstream := chanFromConn(client)
			local := chanFromConn(c)

			go func() {
				for {
					select {
					case b1 := <-upstream:
						if b1 == nil {
							return
						}
						c.Write(b1)
					case b2 := <-local:
						if b2 == nil {
							return
						}
						client.Write(b2)
					}
				}
			}()
		}(c)
	}
}
