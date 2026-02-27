package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main(){       
	for{
	exibeIntroduction()
	readSitesFromFile()
	exibeMenu()

	command := readCommand()
	
	switch command {                        
	case 1:
		fmt.Println("Exibindo monitoramento...")
		sites := readSitesFromFile()
		for _ , site := range sites {
			initMonitoring(site)
			
		}
		fmt.Println("Monitoramento finalizado!")
		fmt.Println()
		continue
	case 2:
		fmt.Println("Exibindo logs...")
		logs:= returnLogs()
		for _ , log := range logs {
			fmt.Println(log)
		}
		continue
	case 0:
		fmt.Println("Saindo do sistema...")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido!")
		os.Exit(-1)
	}
	}
}


func exibeIntroduction(){
	name := "Heitor"
	version := 1.1
	fmt.Println("Olá, ", name, "! Bem-vindo ao Sistema " + getCurrentTime())
	fmt.Println("Versão do Sistema: ", version)
} 

func exibeMenu() {
	fmt.Println("1 - Exibir monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")
}

func readCommand() int {
	fmt.Print("-> ")
	var command int
	fmt.Scan(&command)
	fmt.Println("Comando selecionado: ", command)
	return command
}

func readSitesFromFile() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    reader := bufio.NewReader(arquivo)
    for {
        linha, err := reader.ReadString('\n')
        linha = strings.TrimSpace(linha)

        sites = append(sites, linha)

        if err == io.EOF {
            break
        }

    }

    arquivo.Close()
    return sites
}

func initMonitoring(site string){
	fmt.Println("monitorando {" , site, "} ...")
	res, err := http.Get(site)
	
	if err != nil {
		fmt.Println("Ocorreu um erro ao acessar o site: ", err)
	}

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		saveLog(res.StatusCode, site ,getCurrentTime(), true)
	}else{
		fmt.Println("Site: ", site, "está com problemas!")
		fmt.Println("Status Code: ", res.StatusCode)
		saveLog(res.StatusCode, site ,getCurrentTime(), false)
	}
}

func returnLogs() []string{
	fmt.Println("Exibindo logs...")

	var logs []string
	file , err := os.Open("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler os logs: ", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Fim dos logs!")
			break
		}
		line = strings.TrimSpace(line)
		logs = append(logs, line)
	}
	file.Close()
	return logs
}

func getCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("02/01/2006 15:04:05")
}

func saveLog(statusCode int, site string, timestamp string, status bool){
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	log := fmt.Sprintf("%s = %s => %d : %s\n", timestamp, site, statusCode, strconv.FormatBool(status))
	fmt.Println(log)

	if err != nil {
		fmt.Println("Ocorreu um erro ao salvar o log: ", err)
	}

	file.WriteString(log)

	file.Close()
}

