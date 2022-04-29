package main

import (
	"github.com/miekg/dns"
	"fmt"
)

func main() {

	m := new(dns.Msg)
	m.Id = dns.Id()
	m.RecursionDesired = true
	m.Question = make([]dns.Question, 1)
	m.Question[0] = dns.Question{
		Name: "_dmarc.dina.io.", 
		Qtype: dns.TypeTXT, 
		Qclass: dns.ClassINET }

	c := new(dns.Client)
	in, _, err := c.Exchange(m, "192.168.101.117:53")

	if err != nil {
		fmt.Printf("Erro ao enviar a questão DNS [%s]", err.Error())
	}

	fmt.Println("in.Answer ", in.Answer[0])
}