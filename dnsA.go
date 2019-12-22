package main

import (
	"fmt"
	"os"
	"github.com/miekg/dns"
)

func main() {
	arg := os.Args[1]
	var msg dns.Msg
	fqdn := dns.Fqdn(string(arg))
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}
	for _, answer := range in.Answer {
		if a, ok:= answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
