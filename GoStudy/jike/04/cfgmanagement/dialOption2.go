package main

import "net"

func main() {

}

type dialOptions2 struct {
	dial func(string, string) (net.Conn, error)
}

type DialOption2 func(*dialOptions2)

func Dial2(network, address string, options ...DialOption2) (net.Conn, error) {
	do := dialOptions2{
		dial: net.Dial,
	}

	for _, option := range options {
		option(&do)
	}

	return net.Dial("", "")
}
