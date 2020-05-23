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
	lbBaseURL := "duolingo-leaderboards-prod.duolingo.com"

	fmt.Println("Getting page IDs...")
	pageParseRes, err := duo.ParseMainPage("https://" + baseURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\tLeaderboardID: %s\n\tBasePath: %s\n",
		pageParseRes.LeaderboardID,
		pageParseRes.BasePath)

	duoClient := duo.Client{
		Auth:     httptransport.BearerToken(token),
		Client:   client.New(httptransport.New(baseURL, pageParseRes.BasePath, nil), strfmt.Default),
		ClientLB: client.New(httptransport.New(lbBaseURL, "", nil), strfmt.Default),
	}

	params := users.NewGetUsersParamsWithTimeout(5 * time.Second)
	params.SetUsername(username)

	fmt.Printf("Testing unlogged user data...")
	myUser, err := duoClient.GetUser(username)
	if err != nil {
		panic(err)
	}
	userID := myUser.ID
	fmt.Printf("OK, ID: %d\n", myUser.ID)

	fmt.Printf("Testing leaderboard data...")
	myLb, err := duoClient.GetLeaderboards(pageParseRes.LeaderboardID, userID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Contenst data:")
	fmt.Printf("\tStart %s\n", myLb.Active.Contest.ContestStart)
	fmt.Printf("\tEnd %s\n", myLb.Active.Contest.ContestEnd)
	fmt.Printf("\tNumber of winners %d\n", myLb.Active.Contest.Ruleset.NumWinners)
	fmt.Printf("\tNumber of losers %d\n", myLb.Active.Contest.Ruleset.NumLosers)
}
