package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type hbAccaunts struct {
	Success  bool `json:"success"`
	Accounts []struct {
		ID           string `json:"id"`
		ExtraDetails bool   `json:"extra_details"`
	}
}

func hbAccTaker() []string {

	apiID := "ea91342f6734cc993eb2"
	apiKey := "f3175b6ab1c9c138d1d5&call"

	url := "http://prodportal.kazteleport.kz/admin/api.php?api_id=" + apiID + "&api_key=" + apiKey + "=getAccounts"
	method := "GET"

	client := &http.Client{}

	reqAcc, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	reqAcc.Header.Add("Cookie", "SESSIDb0a7=i65ssikera39g6feuuad2ul2rs")

	resAcc, err := client.Do(reqAcc)
	if err != nil {
		fmt.Println(err)
	}
	defer resAcc.Body.Close()

	body, err := io.ReadAll(resAcc.Body)
	if err != nil {
		fmt.Println(err)
	}
	var hbAcc hbAccaunts
	json.Unmarshal(body, &hbAcc)

	sliceAcc := []string{}
	for i := range hbAcc.Accounts {
		sliceAcc = append(sliceAcc, hbAcc.Accounts[i].ID)

	}
	return sliceAcc
}
