package duo

import (
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/igorskh/go-duolingo/client"
)

var baseURL = "localhost:4010"

var duoClient = Client{
	Auth:      httptransport.BearerToken("token"),
	Client:    client.New(httptransport.New(baseURL, "", []string{"http"}), strfmt.Default),
	ClientLB:  client.New(httptransport.New(baseURL, "", []string{"http"}), strfmt.Default),
	ClientACH: client.New(httptransport.New(baseURL, "", []string{"http"}), strfmt.Default),
}

func TestUser(t *testing.T) {
	_, err := duoClient.GetUser("myuser")
	if err != nil {
		t.Error(err)
	}
}
func TestUser2(t *testing.T) {
	_, err := duoClient.GetUser2(123, []string{"test", "fields"})
	if err != nil {
		t.Error(err)
	}
}
func TestSubscriptions(t *testing.T) {
	_, err := duoClient.GetSubscriptions(123)
	if err != nil {
		t.Error(err)
	}
}
func TestXpSummaries(t *testing.T) {
	_, err := duoClient.GetXpSummaries(123, "utc")
	if err != nil {
		t.Error(err)
	}
}
func TestLeaderboards(t *testing.T) {
	_, err := duoClient.GetLeaderboards("any", 1234)
	if err != nil {
		t.Error(err)
	}
}

func TestAchievements(t *testing.T) {
	_, err := duoClient.GetAchievements(1234, "de", "ru", true, false, false, false)
	if err != nil {
		t.Error(err)
	}
}

func TestShopItems(t *testing.T) {
	_, err := duoClient.GetShopItems()
	if err != nil {
		t.Error(err)
	}
}
