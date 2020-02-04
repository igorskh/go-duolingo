package main

import (
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/igorskh/go-duolingo/client"
	"github.com/igorskh/go-duolingo/duo"
)

func main() {
	token := os.Getenv("DUO_TOKEN")
	username := os.Getenv("DUO_USERNAME")
	baseURL := "www.duolingo.com"

	pageParseRes, err := duo.ParseMainPage("https://" + baseURL)
	if err != nil {
		panic(err)
	}
	duoClient := duo.Client{
		Auth:   httptransport.BearerToken(token),
		Client: client.New(httptransport.New(baseURL, pageParseRes.BasePath, nil), strfmt.Default),
	}

	myUser, err := duoClient.GetUser(username)
	if err != nil {
		panic(err)
	}
	fmt.Println(myUser.ID)
}
