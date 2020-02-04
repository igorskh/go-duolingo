package main

import (
	"fmt"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/igorskh/go-duolingo/client"
	"github.com/igorskh/go-duolingo/client/users"
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

	cl := client.New(httptransport.New(baseURL, pageParseRes.BasePath, nil), strfmt.Default)
	auth := httptransport.BearerToken(token)

	params := users.NewGetUsersParamsWithTimeout(5 * time.Second)
	params.SetUsername(username)

	resp, err := cl.Users.GetUsers(params, auth)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Payload.Users[0].ID)
}
