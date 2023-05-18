package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type CdnAPICep struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  string `json:"complemento"`
}

func main() {
	viaCEP := make(chan ViaCEP)
	cdnApiCEP := make(chan CdnAPICep)

	go func() {
		viaCEP <- getViaCEP("14412009")
	}()

	go func() {
		cdnApiCEP <- getCdnAPICep("1412009")
	}()

	select {
	case msg1 := <-viaCEP:
		fmt.Print("ViaCEP: ")
		fmt.Print(msg1)
	case msg2 := <-cdnApiCEP:
		fmt.Print("CDN API CEP: ")
		fmt.Print(msg2)
	case <-time.After(time.Second * 1):
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

func getCdnAPICep(cep string) CdnAPICep {
	url := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
	res := getCEP(url)
	var data CdnAPICep
	err := json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o unmarshal: %v\n", err)
	}
	return data
}

func getViaCEP(cep string) ViaCEP {
	url := "https://viacep.com.br/ws/" + cep + "/json/"
	res := getCEP(url)
	var data ViaCEP
	err := json.Unmarshal(res, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer o unmarshal: %v\n", err)
	}
	return data
}
