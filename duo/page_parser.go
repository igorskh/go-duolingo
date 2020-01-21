package duo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// PageParseResult main page parsing result
type PageParseResult struct {
	LeaderboardID string
	BasePath      string
}

func getPage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return html
}

// ParseMainPage parses main duolingo page to find vars
func ParseMainPage(url string) PageParseResult {
	mainHTML := getPage(url)

	scriptPathRe := regexp.MustCompile(`<script src="([^\s]+app-.+.js)"><\/script>`)
	scriptPath := string(scriptPathRe.FindSubmatch(mainHTML)[1])
	fmt.Println("Parsing script: " + scriptPath)

	scriptCode := getPage("https:" + scriptPath)

	leaderboardIDRe := regexp.MustCompile(`Le=\"([\S]+)\"`)
	basePathRe := regexp.MustCompile(`Pn=N\(\"\/([\d-]+)\"`)
	leaderboardID := leaderboardIDRe.FindSubmatch(scriptCode)
	basePath := basePathRe.FindSubmatch(scriptCode)

	return PageParseResult{
		LeaderboardID: string(leaderboardID[1]),
		BasePath:      string(basePath[1]),
	}
}
