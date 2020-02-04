# Golang Duolingo Unofficial API Implementation

![](https://github.com/igorskh/go-duolingo/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/igorskh/go-duolingo/branch/master/graph/badge.svg?token=gwaNQFJ6Xi)](https://codecov.io/gh/igorskh/go-duolingo)

## Description

This is an unofficial API client in Go. Client stub is generate with codegen from [go-swagger](https://github.com/go-swagger/go-swagger) from [unofficial Duolingo OpenAPI specification](https://github.com/igorskh/duolingo-api). It can be used to generated code for many other languages. 

Check out the [interactive API documentation](https://duolingo-api.roundeasy.now.sh/).

## What is implemented?
* Authentication
* User data, including courses information and progress
* Subscriptions (friends)
* Leaderboards
* XP Summaries

## Prepare
Most of methods require authentication. One way is to provide token, which you get after login and set your enviromental variables:
```bash
export DUO_TOKEN=<your_token_here>
export DUO_USERNAME=<your_username_here>
```

Read them in Go:
```go
token := os.Getenv("DUO_TOKEN")
username := os.Getenv("DUO_USERNAME")
```

Advantage of this approach that you do not need to store your password anywhere. Also, token is stored only for the duration of the terminal session.

## Usage
Basic example [example2/main.go](example2/main.go) shows how to retrieve the user information and output userID by given username.
```go
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

```


## High-level API Access
Package `github.com/igorskh/go-duolingo/duo` is an additional layer of abstraction which simplifies access to the API functions.

Example in [examples/main.go](example1/main.go) shows the process of retrieving userID, but using high-level client.
```go
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
	lbBaseURL := "duolingo-leaderboards-prod.duolingo.com"

	pageParseRes, err := duo.ParseMainPage("https://" + baseURL)
	if err != nil {
		panic(err)
	}
	duoClient := duo.Client{
		Auth:     httptransport.BearerToken(token),
		Client:   client.New(httptransport.New(baseURL, pageParseRes.BasePath, nil), strfmt.Default),
		ClientLB: client.New(httptransport.New(lbBaseURL, "", nil), strfmt.Default),
	}

	myUser, err := duoClient.GetUser(username)
	if err != nil {
		panic(err)
	}
	fmt.Println(myUser.ID)
}
```

## Serialising to JSON
```go
bytes, err := json.Marshal(obj)
if err != nil {
    fmt.Println("Error while serialising", obj)
}
```

Variable `bytes` contains bytes array of JSON string and can be easily converted to `string` as `string(bytes)`.

```go
fmt.Println(string(bytes))
```