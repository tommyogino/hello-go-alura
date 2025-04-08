package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 5

func main() {
	exibeIntroducao()
	for{
		exibeMenu()
		comando := lerComando()
		switch comando {
			case 1:
				iniciarMonitoramento()
			case 2:
				fmt.Println("showing logs...")
				registraLog("site-teste", false)
			case 3:
				fmt.Println("exiting program...")
				os.Exit(0)
			default:
				fmt.Println("invalid command!")
		}
	}
}

func exibeIntroducao(){
	nome  := "Tommy"
	versao := 1.1
	fmt.Println("hello", nome)
	fmt.Println("program version:", versao)
	fmt.Println()
}

func exibeMenu(){
	fmt.Println("1 - start monitoring")
	fmt.Println("2 - show logs")
	fmt.Println("3 - exit program")
	fmt.Println()
}

func lerComando () int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func iniciarMonitoramento(){
	fmt.Println()
	fmt.Println("monitoring...")
	fmt.Println()
	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++{

		for i, site := range sites{
			fmt.Println(i + 1,"-", site)
			testaSite(site)
		
		}
	time.Sleep(delay * time.Second)
	fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string){
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("error:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("site:", site, "is online")
		registraLog(site, true)
	} else {
		fmt.Println("site:", site, "is down. statusCode:", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() [] string {
	var sites [] string
	arquivo, err := os.Open("sitesMonitorados.txt")

	if err != nil{
		fmt.Println("error:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		sites = append(sites, linha)
		
		if err == io.EOF {
			break	
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt",os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("error:", err)

	}
	fmt.Println(arquivo, site, status)
}