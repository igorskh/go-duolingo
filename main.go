package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/igorskh/go-duolingo/client"
	"github.com/igorskh/go-duolingo/client/users"
	"github.com/igorskh/go-duolingo/duo"
	"github.com/igorskh/go-duolingo/models"
)

func getUser(cl *client.DuolingoUnofficial, auth *runtime.ClientAuthInfoWriter, username string) (*models.User, error) {
	params := users.NewGetUsersParamsWithTimeout(5 * time.Second)
	params.SetUsername(username)

	resp, err := cl.Users.GetUsers(params, *auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Users[0], nil
}

func getSubscriptions(cl *client.DuolingoUnofficial, auth *runtime.ClientAuthInfoWriter, userID int64) (*models.SubscriptionList, error) {
	params := users.NewGetSubscriptionsParamsWithTimeout(5 * time.Second)
	params.SetUserID(userID)

	resp, err := cl.Users.GetSubscriptions(params, *auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func main() {
	token := os.Getenv("DUO_TOKEN")
	username := os.Getenv("DUO_USERNAME")
	baseURL := "www.duolingo.com"

	pageParseRes, err := duo.ParseMainPage("https://" + baseURL)
	if err != nil {
		panic(err)
	}

	auth := httptransport.BearerToken(token)
	cl := client.New(httptransport.New(baseURL, pageParseRes.BasePath, nil), strfmt.Default)

	myUser, err := getUser(cl, &auth, username)
	if err != nil {
		panic(err)
	}

	subs, err := getSubscriptions(cl, &auth, myUser.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(myUser.ID)
	fmt.Println(myUser.Courses[0].Crowns)
	fmt.Println(subs.Subscriptions[0])
}
