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

const monitoramentos = 3
const delay = 5

func main() {
	introdução()
	for {
		menu()
		comando := leComando()

		switch comando {
		case 1:
			monitoramento()

		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLog()

		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)

		default:
			fmt.Println("Comando invalido")
			os.Exit(-1)
		}
	}
}

func introdução() {
	version := 1.0
	fmt.Println("Esse programa esta na versão ", version)
}

func menu() {
	fmt.Println("-------------|MENU|------------")
	fmt.Println("| 1- Inicia monitoramento     |")
	fmt.Println("| 2- Exibe log's              |")
	fmt.Println("| 0- Sai do programa          |")
	fmt.Println("-------------------------------")
}

func monitoramento() {

	sites := leArquivo()

	for i := 0; i < monitoramentos; i++ {
		fmt.Println("")
		fmt.Println("Teste: ", i+1)
		for i, site := range sites {
			fmt.Println("Testando site ", i, ":", site)
			testasite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testasite(site string) {

	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if response.StatusCode == 200 {
		fmt.Println("Site: ", site, " esta OK!")
		fmt.Println("")
		registraLog(site, true)
	} else {
		fmt.Println("🚨 Site: ", site, "esta com algum problema. Status code: ", response.StatusCode)
		fmt.Println("")
		registraLog(site, false)
	}
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func leArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err_arquivo := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if err_arquivo == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	} //-----------------------------------dd/mm/yyyy hh:mm:ss
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "- online: " + strconv.FormatBool(status) + "\n")
}

func imprimeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
