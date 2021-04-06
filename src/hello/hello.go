package main

import (
	"fmt"
	"net/http"
	"os"
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
		fmt.Println("Exibindo Logs...")
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
	sites := []string{
		"https://www.alura.com.br",
		"https://www.google.com.br",
		"https://httpstat.us/422",
	}

	for _, site := range sites {
		testaSite(site)
		time.Sleep(5 * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode > 199 && resp.StatusCode < 300 {
		fmt.Println("Site:", site, "status:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "status:", resp.StatusCode)
	}
}
