package apivercel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCep(cep string) (response Response) {
	url := fmt.Sprintf("https://cep-api.vercel.app/api/%s", cep)

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
		fmt.Println(*responseVercel.Info.Message)
		return
	}

	return Response{
		Cep:      responseVercel.Info.Cep,
		Uf:       responseVercel.Info.State,
		City:     responseVercel.Info.City,
		District: responseVercel.Info.District,
		Address:  responseVercel.Info.Address,
	}
}