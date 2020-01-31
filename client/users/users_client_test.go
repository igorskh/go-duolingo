package users

import (
	"testing"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	defaultTimeout = 5 * time.Second
	baseURL        = "localhost:4010"
)

var (
	cl   = New(httptransport.New(baseURL, "", []string{"http"}), strfmt.Default)
	auth = httptransport.BearerToken("token")
)

func TestUser(t *testing.T) {
	params := NewGetUsersParamsWithTimeout(defaultTimeout)
	params.SetUsername("username")

	_, err := cl.GetUsers(params, auth)
	if err != nil {
		t.Error(err)
	}
}
func TestUser2(t *testing.T) {
	params := NewGetUserParamsWithTimeout(defaultTimeout)
	params.SetUserID(123)
	params.SetFields([]string{"test", "fields"})

	_, err := cl.GetUser(params, auth)
	if err != nil {
		t.Error(err)
	}
}
func TestSubscriptions(t *testing.T) {
	params := NewGetSubscriptionsParamsWithTimeout(defaultTimeout)
	params.SetUserID(123)

	_, err := cl.GetSubscriptions(params, auth)
	if err != nil {
		t.Error(err)
	}
}
func TestXpSummaries(t *testing.T) {
	params := NewGetXpSummariesParamsWithTimeout(defaultTimeout)
	params.SetUserID(123)
	params.SetTimezone("UTC")

	_, err := cl.GetXpSummaries(params, auth)
	if err != nil {
		t.Error(err)
	}
}
