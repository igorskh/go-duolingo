# Golang Duolingo Unofficial API Implementation

![](https://github.com/igorskh/go-duolingo/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/igorskh/go-duolingo/branch/master/graph/badge.svg?token=gwaNQFJ6Xi)](https://codecov.io/gh/igorskh/go-duolingo)

## What is implemented?
* Authentication
* User data, including courses information and progress
* Subscriptions (friends)
* Leaderboards
* XP Summaries

## Prepare
Set your enviromental variables:
```bash
export DUO_TOKEN=<your_token_here>
export DUO_USERNAME=<your_username_here>
```

Read them in Go:
```go
token := os.Getenv("DUO_TOKEN")
username := os.Getenv("DUO_USERNAME")
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
