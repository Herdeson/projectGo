package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	//fmt.Scanf("%d", &comando)

	// if comando == 1 {
	// 	fmt.Println("Iniciando o monitoramento..")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo os logs")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do Programa")
	// } else {
	// 	fmt.Println("Nenhuma opção não fornecida")
	// }

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
	fmt.Println("Iniciando o monitoramento..")
	site := "https://www.alura.com.br"
	resp, error := http.Get(site)
	fmt.Println(resp)
	fmt.Println(error)

}
