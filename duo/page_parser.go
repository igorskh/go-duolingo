package duo

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// PageParseResult main page parsing result
type PageParseResult struct {
	LeaderboardID string
	BasePath      string
}

func getPage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return html, nil
}

// ParseMainPage parses main duolingo page to find vars
func ParseMainPage(url string) (*PageParseResult, error) {
	mainHTML, err := getPage(url)
	if err != nil {
		return nil, err
	}

	scriptPathRe := regexp.MustCompile(`<script src="([^\s]+app-.+.js)"><\/script>`)
	scriptPath := string(scriptPathRe.FindSubmatch(mainHTML)[1])
	scriptCode, err := getPage("https:" + scriptPath)
	if err != nil {
		return nil, err
	}

	leaderboardIDRe := regexp.MustCompile(`Le=\"([\S]+)\"`)
	basePathRe := regexp.MustCompile(`Pn=N\(\"\/([\d-]+)\"`)
	leaderboardID := leaderboardIDRe.FindSubmatch(scriptCode)
	basePath := basePathRe.FindSubmatch(scriptCode)

	return &PageParseResult{
		LeaderboardID: string(leaderboardID[1]),
		BasePath:      string(basePath[1]),
	}, nil
}
