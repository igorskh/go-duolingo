package main

import (
	"flag"
	"fmt"
	"log"
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
	baseURLPtr := flag.String(
		"url",
		"www.duolingo.com",
		"Duolingo main API URL")
	lbBaseURLPtr := flag.String(
		"lb-url",
		"duolingo-leaderboards-prod.duolingo.com",
		"Duolingo leaderboards API URL")
	achBaseURLPtr := flag.String(
		"ach-url",
		"duolingo-achievements-prod.duolingo.com",
		"Duolingo achievements API URL")
	flag.Parse()

	if usernamePtr == nil || *usernamePtr == "" {
		log.Fatalln("Username is required")
	}
	token := *tokenPtr
	userID := *userIDPtr
	username := *usernamePtr
	baseURL := *baseURLPtr
	lbBaseURL := *lbBaseURLPtr
	achBaseURL := *achBaseURLPtr

	fmt.Println("Getting page IDs...")
	pageParseRes, err := duo.ParseMainPage("https://" + baseURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\tLeaderboardID: %s\n\tBasePath: %s\n",
		pageParseRes.LeaderboardID,
		pageParseRes.BasePath)

	duoClient := duo.Client{
		Auth:      httptransport.BearerToken(token),
		Client:    client.New(httptransport.New(baseURL, pageParseRes.BasePath, nil), strfmt.Default),
		ClientLB:  client.New(httptransport.New(lbBaseURL, "", nil), strfmt.Default),
		ClientACH: client.New(httptransport.New(achBaseURL, "", nil), strfmt.Default),
	}

	fmt.Printf("Testing unlogged user data...")
	myUser, err := duoClient.GetUser(username)
	if err != nil {
		panic(err)
	}
	if userID == 0 {
		userID = myUser.ID
	}
	fmt.Printf("OK, ID: %d\n", myUser.ID)

	fmt.Printf("Testing logged in user data...")
	fields := []string{"id", "courses", "fromLanguage", "learningLanguage"}
	myUser, err = duoClient.GetUser2(userID, fields)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OK, first course crowns: %d\n", myUser.Courses[0].Crowns)

	fmt.Printf("Testing leaderboard data...")
	myLb, err := duoClient.GetLeaderboards(pageParseRes.LeaderboardID, userID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OK, tier: %d\n", myLb.Tier)

	fmt.Printf("Testing subscriptions...")
	subs, err := duoClient.GetSubscriptions(userID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OK, firstID: %d\n", subs.Subscriptions[0].ID)

	fmt.Printf("Testing achievements...")
	achs, err := duoClient.GetAchievements(
		userID,
		myUser.FromLanguage,
		myUser.LearningLanguage,
		myUser.HasPlus,
		false,
		true,
		false)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OK, firstID: %s\n", achs.Achievements[0].Name)

	fmt.Printf("Testing shop items...")
	shopItems, err := duoClient.GetShopItems()
	if err != nil {
		panic(err)
	}
	fmt.Printf("OK, shop item 0: %s\n", shopItems.ShopItems[0].Name)
}
