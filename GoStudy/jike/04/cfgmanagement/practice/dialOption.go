package main

import "fmt"

func main() {
	cfg := NewRdsCfg("", "", "", dialTimeOut("1s"))

	fmt.Println(cfg.timeout)
}

type dialOptions struct {
	dial    func()
	timeout string
}

type DialOption struct {
	f func(*dialOptions)
}

type rdsCfg struct {
	address string
	ip      string
	port    string
	timeout string
}

func NewRdsCfg(address, ip, port string, dialOption ...DialOption) *rdsCfg {
	do := &dialOptions{dial: func() {

	}}
	for _, d := range dialOption {
		d.f(do)
	}

	return &rdsCfg{
		address: address,
		ip:      ip,
		port:    port,
		timeout: do.timeout,
	}
}

func dialTimeOut(tm string) DialOption {
	return DialOption{f: func(options *dialOptions) {
		options.timeout = tm
	}}
}
