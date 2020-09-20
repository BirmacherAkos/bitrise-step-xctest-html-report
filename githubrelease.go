package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bitrise-io/go-utils/log"
)

// GithubRelease returned from the Github API
type GithubRelease struct {
	TagName string `json:"tag_name"`
}

// https://api.github.com/repos/TitouanVanBelle/XCTestHTMLReport/releases/latest

// Returns the information of the latest release from the Github API for the provided repository
// More: https://docs.github.com/en/rest/reference/repos#releases
func latestGithubRelease(githubOrg, githubRepository string) (GithubRelease, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", githubOrg, githubRepository)
	resp, err := http.Get(url)
	if err != nil {
		return GithubRelease{}, fmt.Errorf("failed to call the %s, error: %v", url, err)
	}

	log.Debugf("Response status: %s", resp.Status)
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Warnf("Failed to close response body, error: %v", cerr)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GithubRelease{}, fmt.Errorf("failed to read response body from the %s, error: %v", url, err)
	}
	var githubRelease GithubRelease
	json.Unmarshal([]byte(body), &githubRelease)

	return githubRelease, nil
}
