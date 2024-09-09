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

const monitoramento = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		fmt.Print("\n")
		fmt.Print("\n")
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()

		case 2:

			fmt.Println("Exibindo Logs...")
			imprimeLogs()

		case 3:
			fmt.Println("Saindo do programa...")
			os.Exit(0)

		default:
			fmt.Println("Nao conheco este comando.")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Cliente"
	versao := 1.1

	fmt.Println("Ola, sr.", nome)
	fmt.Println("Este programa esta na versao", versao)
	fmt.Print("\n")
}

func exibeMenu() {
	fmt.Println("------ MENU ------")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")
	fmt.Println("-------------------")
	fmt.Println("\n")
}

func leComando() int {
	var comandoLido int
	fmt.Print("Resposta:")
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("Comando escolhido foi o: ", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site ", i, ": ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Minute)
		fmt.Print("\n")
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registrarLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registrarLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n') //barra simples pois recebe byte
		linha = strings.TrimSpace(linha)      //tira espaco desnecessario

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
