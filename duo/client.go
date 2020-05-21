package duo

import (
	"time"

	"github.com/go-openapi/runtime"
	"github.com/igorskh/go-duolingo/client"
	"github.com/igorskh/go-duolingo/client/achievements"
	"github.com/igorskh/go-duolingo/client/leaderboards"
	"github.com/igorskh/go-duolingo/client/shop"
	"github.com/igorskh/go-duolingo/client/users"
	"github.com/igorskh/go-duolingo/models"
)

// Client client for accessing Duolingo API
type Client struct {
	Client    *client.DuolingoUnofficial
	ClientLB  *client.DuolingoUnofficial
	ClientACH *client.DuolingoUnofficial
	Auth      runtime.ClientAuthInfoWriter
}

// GetUser gets user info without authentication
func (c Client) GetUser(username string) (*models.User, error) {
	params := users.NewGetUsersParamsWithTimeout(5 * time.Second)
	params.SetUsername(username)

	resp, err := c.Client.Users.GetUsers(params, c.Auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Users[0], nil
}

// GetUser2 gets user info with authentication
func (c Client) GetUser2(userID int64, fields []string) (*models.User, error) {
	params := users.NewGetUserParamsWithTimeout(5 * time.Second)
	params.SetUserID(userID)
	params.SetFields(fields)

	resp, err := c.Client.Users.GetUser(params, c.Auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetSubscriptions gets user subscriptions (friends)
func (c Client) GetSubscriptions(userID int64) (*models.SubscriptionList, error) {
	params := users.NewGetSubscriptionsParamsWithTimeout(5 * time.Second)
	params.SetUserID(userID)

	resp, err := c.Client.Users.GetSubscriptions(params, c.Auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetLeaderboards get user leaderboard information
func (c Client) GetLeaderboards(lbID string, userID int64) (*models.Leaderdoard, error) {
	params := leaderboards.NewGetLeaderboardParamsWithTimeout(5 * time.Second)
	params.SetUserID(userID)
	params.SetLeaderboardID(lbID)

	resp, err := c.ClientLB.Leaderboards.GetLeaderboard(params, c.Auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetShopItems get available shop items
func (c Client) GetShopItems() (*models.ShopItemList, error) {
	params := shop.NewGetShopItemListParamsWithTimeout(5 * time.Second)

	resp, err := c.Client.Shop.GetShopItemList(params, c.Auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

// GetXpSummaries get user XP summaries
func (c Client) GetXpSummaries(userID int64, timezone string) ([]*models.XpSummary, error) {
	params := users.NewGetXpSummariesParamsWithTimeout(5 * time.Second)
	params.SetUserID(userID)
	params.SetTimezone(timezone)

	resp, err := c.Client.Users.GetXpSummaries(params, c.Auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload.Summaries, nil
}

// GetAchievements gets user achievements
func (c Client) GetAchievements(
	userID int64,
	fromLanguage string,
	learningLanguage string,
	hasPlus bool,
	isAgeRestricted bool,
	isProfilePublic bool,
	isSchools bool) (*models.Achievements, error) {
	params := achievements.NewGetAchievementsParamsWithTimeout(5 * time.Second)
	params.SetUserID(userID)
	params.SetFromLanguage(fromLanguage)
	params.SetLearningLanguage(learningLanguage)

	falseVal := int64(0)
	trueVal := int64(1)
	if hasPlus {
		params.SetHasPlus(&trueVal)
	} else {
		params.SetHasPlus(&falseVal)
	}
	if isAgeRestricted {
		params.SetIsAgeRestricted(&trueVal)
	} else {
		params.SetIsAgeRestricted(&falseVal)
	}
	if isProfilePublic {
		params.SetIsProfilePublic(&trueVal)
	} else {
		params.SetIsProfilePublic(&falseVal)
	}
	if isSchools {
		params.SetIsSchools(&trueVal)
	} else {
		params.SetIsSchools(&falseVal)
	}

	resp, err := c.ClientACH.Achievements.GetAchievements(params, c.Auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
