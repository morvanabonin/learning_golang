package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Teste de limpeza")

	//sDomain := "mail.mail.mail.mail.mail.mail.mail.cli.esrvvenancio.com.br."
	sDomain := "mail.mail.mail.mail.mta-sts.5cde831a.cli.dnzpark.com.br."

	if strings.Contains(sDomain, "mail.") {
		rx := regexp.MustCompile(`mail.`)
		matches := rx.FindAllString(sDomain, -1)
		fmt.Println(matches)
		mailString := strings.Join(matches, "")
		//fmt.Println(mailString)
		fmt.Println(strings.Replace(sDomain, mailString, "mail.", -1))
	}

}
