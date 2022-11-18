package main

import (
	"fmt"
)

func main() {
	var nome string = "Douglas"
	var versao = 1.1

	fmt.Println("Olá sr", nome)
	fmt.Println("Progama na versão:", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
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
		fmt.Println("Iniciando o monitoramento..")
	case 2:
		fmt.Println("Exibindo os logs")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Nenhuma opção conhecida")
	}

}
