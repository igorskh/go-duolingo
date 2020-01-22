package main

import (
	"flag"
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/igorskh/go-duolingo/client"
	"github.com/igorskh/go-duolingo/duo"
)

func main() {
	tokenPtr := flag.String("token", os.Getenv("DUO_TOKEN"), "Duolingo token")
	usernamePtr := flag.String("username", os.Getenv("DUO_USERNAME"), "Duolingo userID")
	userIDPtr := flag.Int64("userid", 0, "Duolingo username")
	baseURLPtr := flag.String("url", "www.duolingo.com", "Duolingo main API URL")
	lbBaseURLPtr := flag.String("lb-url", "duolingo-leaderboards-prod.duolingo.com", "Duolingo leaderboards API URL")
	flag.Parse()

	token := *tokenPtr
	userID := *userIDPtr
	username := *usernamePtr
	baseURL := *baseURLPtr
	lbBaseURL := *lbBaseURLPtr

	pageParseRes, err := duo.ParseMainPage("https://" + baseURL)
	if err != nil {
		panic(err)
	}
	duoClient := duo.Client{
		Auth:     httptransport.BearerToken(token),
		Client:   client.New(httptransport.New(baseURL, pageParseRes.BasePath, nil), strfmt.Default),
		ClientLB: client.New(httptransport.New(lbBaseURL, "", nil), strfmt.Default),
	}

	if userID == 0 {
		myUser, err := duoClient.GetUser(username)
		if err != nil {
			panic(err)
		}
		userID = myUser.ID
	}

	fields := []string{"id", "courses"}
	myUser, err := duoClient.GetUser2(userID, fields)
	if err != nil {
		panic(err)
	}

	myLb, err := duoClient.GetLeaderboards(pageParseRes.LeaderboardID, userID)
	if err != nil {
		panic(err)
	}

	subs, err := duoClient.GetSubscriptions(userID)
	if err != nil {
		panic(err)
	}

	fmt.Println(myUser.ID)
	fmt.Println(myUser.Courses[0].Crowns)
	fmt.Println(subs.Subscriptions[0])
	fmt.Println(myLb.Tier)
}
