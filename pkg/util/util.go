//
// Copyright 2021-2022 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
	transportHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
)

func SanitizeName(name string) string {
	sanitizedName := strings.ToLower(strings.Replace(strings.Replace(name, " ", "-", -1), "'", "", -1))
	if len(sanitizedName) > 50 {
		sanitizedName = sanitizedName[0:50]
	}

	return sanitizedName
}

// IsExist returns whether the given file or directory exists
func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ConvertGitHubURL converts a git url to its raw format
// adapted from https://github.com/redhat-developer/odo/blob/e63773cc156ade6174a533535cbaa0c79506ffdb/pkg/catalog/catalog.go#L72
func ConvertGitHubURL(URL string) (string, error) {
	// If the URL ends with .git, remove it
	// The regex will only instances of '.git' if it is at the end of the given string
	reg := regexp.MustCompile(".git$")
	URL = reg.ReplaceAllString(URL, "")

	url, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	if strings.Contains(url.Host, "github") && !strings.Contains(url.Host, "raw") {
		// Convert path part of the URL
		URLSlice := strings.Split(URL, "/")
		if len(URLSlice) > 2 && URLSlice[len(URLSlice)-2] == "tree" {
			// GitHub raw URL doesn't have "tree" structure in the URL, need to remove it
			URL = strings.Replace(URL, "/tree", "", 1)
		} else {
			// Add "main" branch for GitHub raw URL by default if branch is not specified
			URL = URL + "/main"
		}

		// Convert host part of the URL
		if url.Host == "github.com" {
			URL = strings.Replace(URL, "github.com", "raw.githubusercontent.com", 1)
		}
	}

	return URL, nil
}

// CurlEndpoint curls the endpoint and returns the response or an error if the response is a non-200 status
func CurlEndpoint(endpoint string) ([]byte, error) {
	var respBytes []byte
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		respBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return respBytes, nil
	}

	return nil, fmt.Errorf("received a non-200 status when curling %s", endpoint)
}

// CloneRepo clones the repoURL to clonePath
func CloneRepo(clonePath, repoURL string, token string) error {
	// Set up the Clone options
	cloneOpts := &git.CloneOptions{
		URL: repoURL,
	}

	// If a token was passed in, configure token auth for the git client
	if token != "" {
		cloneOpts.Auth = &transportHttp.BasicAuth{
			Username: "token",
			Password: token,
		}
	}
	// Clone the repo
	_, err := git.PlainClone(clonePath, false, cloneOpts)
	if err != nil {
		return err
	}

	return nil
}
