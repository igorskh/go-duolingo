package main

import (
	"fmt"
	"os"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/igorskh/go-duolingo/client"
	"github.com/igorskh/go-duolingo/client/operations"
	"github.com/igorskh/go-duolingo/duo"
)

func main() {
	token := os.Getenv("DUO_TOKEN")
	username := os.Getenv("DUO_USERNAME")
	baseURL := "www.duolingo.com"

	pageParseRes := duo.ParseMainPage("https://" + baseURL)

	cl := client.New(httptransport.New(baseURL, pageParseRes.BasePath, nil), strfmt.Default)
	params := operations.NewGetUsersParams()
	params.SetTimeout(10 * time.Second)
	params.SetUsername(username)

	writer := httptransport.BearerToken(token)
	resp, err := cl.Operations.GetUsers(params, writer)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp.Payload.Users[0].XpGains)
}
