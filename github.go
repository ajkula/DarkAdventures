// "Package main" is the namespace declaration
package main

// importing standard libraries
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// constants
const (
	githugApiURL = "https://api.github.com/repos/ajkula/DarkAdventures/tags"
)

// User struct represents the JSON data from GitHub API: https://api.github.com/repos/ajkula/DarkAdventures/tags

type Repo struct {
	Name       string `json:"name"`
	ZipballUrl string `json:"zipball_url"`
	TarballUrl string `json:"tarball_url"`
	Commit     Commit `json:"commit"`
	node_id    string `json:"node_id"`
}

type Commit struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

var Repos []*Repo

func getRepos() *Repo {
	resp, err := http.Get(githugApiURL)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	json.Unmarshal(body, &Repos)
	return Repos[0]
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error retrieving Game last version: %s\n", e)
	}
}
