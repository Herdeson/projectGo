package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 3

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo os logs")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}

	}
}
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func exibeIntroducao() {
	var nome string = "Douglas"
	var versao = 1.1

	fmt.Println("Olá sr", nome)
	fmt.Println("Progama na versão:", versao)
}

func iniciarMonitoramento() {
	fmt.Println("")
	fmt.Println("Iniciando o monitoramento..")
	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i <= monitoramentos; i++ {
		for _, item := range sites {
			testaSite(item)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)

	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "Está com problemas")
	}
}
