package main

import "net"

func main() {

}

type dialOptions struct {
	dial func(string, string) (net.Conn, error)
}

type DialOption struct {
	f func(*dialOptions)
}

func Dial(network, address string, options ...DialOption) (net.Conn, error) {
	do := dialOptions{
		dial: net.Dial,
	}

	for _, option := range options {
		option.f(&do)
	}

	return net.Dial("", "")
}
