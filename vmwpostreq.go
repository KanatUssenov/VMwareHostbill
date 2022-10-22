package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type vmwSession struct {
	XMLName    xml.Name `xml:"Session"`
	Text       string   `xml:",chardata"`
	Xmlns      string   `xml:"xmlns,attr"`
	Vmext      string   `xml:"vmext,attr"`
	Ovf        string   `xml:"ovf,attr"`
	Vssd       string   `xml:"vssd,attr"`
	Common     string   `xml:"common,attr"`
	Rasd       string   `xml:"rasd,attr"`
	Vmw        string   `xml:"vmw,attr"`
	Ovfenv     string   `xml:"ovfenv,attr"`
	Ns9        string   `xml:"ns9,attr"`
	User       string   `xml:"user,attr"`
	Org        string   `xml:"org,attr"`
	UserId     string   `xml:"userId,attr"`
	Roles      string   `xml:"roles,attr"`
	LocationId string   `xml:"locationId,attr"`
	Href       string   `xml:"href,attr"`
	Type       string   `xml:"type,attr"`
	Link       []struct {
		Text string `xml:",chardata"`
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
		Type string `xml:"type,attr"`
		Name string `xml:"name,attr"`
	} `xml:"Link"`
	AuthorizedLocations struct {
		Text     string `xml:",chardata"`
		Location struct {
			Text            string `xml:",chardata"`
			LocationId      string `xml:"LocationId"`
			SiteName        string `xml:"SiteName"`
			OrgName         string `xml:"OrgName"`
			RestApiEndpoint string `xml:"RestApiEndpoint"`
			UIEndpoint      string `xml:"UIEndpoint"`
			AuthContext     string `xml:"AuthContext"`
			ApiVersion      string `xml:"ApiVersion"`
		} `xml:"Location"`
	} `xml:"AuthorizedLocations"`
}

func vmwposreq() {

	url := "https://vcloud.kazteleport.kz/api/sessions"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/*;version=36.2")
	req.Header.Add("Authorization", "Basic ZHJvZ2EueXVAc3lzdGVtOlMmeDRQbGVhc3VyZSFAIw==")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var vmwSessionInf vmwSession
	xml.Unmarshal(body, &vmwSessionInf)

	fmt.Println(vmwSessionInf)
}
