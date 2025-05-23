package main

import (
	"curso-go-extensive-desafio-multithreading/src/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	viaCEP := make(chan entities.ViaCEP) //inicializando canal
	cdnApiCEP := make(chan entities.CdnAPICep)

	go func() {
		if os.Getenv("SET_CDNAPI_TIMEOUT") == "true" {
			time.Sleep(time.Second * 10)
		}
		cdnApiCEP <- getCdnAPICep("14412-009")
	}()

	go func() {
		if os.Getenv("SET_VIACEP_TIMEOUT") == "true" {
			time.Sleep(time.Second * 10)
		}
		fmt.Println(os.Getenv("SET_VIACEP_TIMEOUT"))
		viaCEP <- getViaCEP("14412-009")
	}()

	select {
	case viaCepResponse := <-viaCEP:
		fmt.Print("ViaCEP: ")
		fmt.Print(viaCepResponse)
	case cdnApiCepResponse := <-cdnApiCEP:
		fmt.Print("CDN API CEP: ")
		fmt.Print(cdnApiCepResponse)
	case <-time.After(time.Second * 20):
		fmt.Println("timeout")
	}
}

func getCEP(url string) []byte {
	req, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer a leitura do response da api: %v\n", err)
	}
	return res
}

func getCdnAPICep(cep string) entities.CdnAPICep {
	fmt.Print("Starting getCdnAPICep")
	url := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
	res := getCEP(url)
	var data entities.CdnAPICep
	err := json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o unmarshal: %v\n", err)
	}
	return data
}

func getViaCEP(cep string) entities.ViaCEP {
	fmt.Print("Starting getCdnAPICep")
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	res := getCEP(url)
	var data entities.ViaCEP
	err := json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o unmarshal: %v\n", err)
	}
	return data
}
