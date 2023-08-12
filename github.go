package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	githubApiURL = "https://api.github.com/repos/ajkula/DarkAdventures/tags"
)

type Repo struct {
	Name       string `json:"name"`
	ZipballUrl string `json:"zipball_url"`
	TarballUrl string `json:"tarball_url"`
	Commit     Commit `json:"commit"`
	NodeID     string `json:"node_id"`
}

type Commit struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

var repos []*Repo

func getRepos() *Repo {
	resp, err := http.Get(githubApiURL)
	if err != nil {
		check(err)
		return &Repo{Name: releaseVersion}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	json.Unmarshal(body, &repos)
	if len(repos) == 0 || repos[0] == nil {
		return &Repo{Name: releaseVersion}
	}
	return repos[0]
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error retrieving Game last version: %s\n", e)
	}
}
