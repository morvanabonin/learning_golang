package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Teste de limpeza")

	qName := "mail.mail.mail.mail.mail.mail.mail.cli.esrvvenancio.com.br."
	//qName := "mail.mail.mail.mail.mta-sts.5cde831a.cli.dnzpark.com.br."

	if strings.Contains(qName, "mail.") {
		qName = cleanMailDomain(qName)
		fmt.Println(qName)
	}

}

// cleanMailDomain
// Função criada para limpar os domínios que vem com vários "mail."
func cleanMailDomain(domain string) string {
	// regex apenas para passar o padrão mail,
	// nosso interesse aqui é retornar todos
	rx := regexp.MustCompile(`mail.`)
	matches := rx.FindAllString(domain, -1)

	// queremos a string inteira que veio em formato de slice
	mailString := strings.Join(matches, "")

	// com a string já obtendo a quantidade total de `mail`,
	// nós substituímos tudo por apenas um `mail`
	domainClean := strings.Replace(domain, mailString, "mail.", -1)

	return domainClean
}
