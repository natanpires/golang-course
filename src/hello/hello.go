package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		checarComando(leComando())
	}

}

func checarComando(comando int) {
	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		imprimeLog()
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for _, site := range sites {
		testaSite(site)
		time.Sleep(2 * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode > 199 && resp.StatusCode < 300 {
		fmt.Println("\n|", site, "\n -StatusCode:", resp.StatusCode, "\n -Status: success")
		registraLog(site, true)
	} else {
		fmt.Println("\n|", site, "\n -StatusCode:", resp.StatusCode, "\n -Status: error")
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile(
		"log.txt",
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0666,
	)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(
		time.Now().Format("02/01/2006 15:04:05") +
			" - " +
			site +
			" - online: " +
			strconv.FormatBool(status) + "\n",
	)
	arquivo.Close()
}

func imprimeLog() {
	fmt.Println("Exibindo Logs...")

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
