package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"path/filepath"
	"strconv"

	
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
			imprimeLogs()
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
	sites := leSitesArquivo()

	for i := 0; i <= monitoramentos; i++ {
		for _, item := range sites {
			testaSite(item)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)

	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "Está com problemas")
		registraLog(site, false)
	}
}

func leSitesArquivo() []string {
	var sites []string
	file := "sites.txt"
	full_path, err := filepath.Abs(file)
	if err != nil {
		fmt.Println("Apresentou o seguinte erro:", err)
	}
	arquivo, err := os.Open(full_path) // Abrindo arquivo de forma tradicional

	if err != nil {
		fmt.Println("Apresentou o seguinte erro:", err)
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
	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR | os.O_APPEND, 0666) // Se nao tiver arquivo e criado

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }
	

    arquivo.WriteString( time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv. FormatBool(status)+ "\n")

    arquivo.Close()
}

