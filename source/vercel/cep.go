package apivercel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"locus/config"
	"locus/source"
	"log"
	"net/http"
)

const (
	SourceApi = "vercel"
)

// GetCep return CEP info using => https://cep-api.vercel.app/api
func GetCep(findCep string, messages chan source.Response) {
	url := fmt.Sprintf("https://cep-api.vercel.app/api/%s", findCep)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	responseVercel := ResponseVercel{}
	jsonErr := json.Unmarshal(body, &responseVercel)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	if responseVercel.Info.Status != nil {
		fmt.Println(config.ColorRed, fmt.Sprintf("CEP %s not found!", findCep))
		messages <- source.Response{}
	}

	messages <- source.Response{
		Cep:       responseVercel.Info.Cep,
		Uf:        responseVercel.Info.State,
		City:      responseVercel.Info.City,
		District:  responseVercel.Info.District,
		Address:   responseVercel.Info.Address,
		SourceApi: SourceApi,
	}
}