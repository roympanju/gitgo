package main

import (
	"encoding/json"
	//"time"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiURL = "https://api.github.com"
	userEndpoint = "/users/"
)

type User struct {
	Login		string		`json:"login"`
	ID		int		`json:"id"`
	Html_Url	string		`json:"html_url"`
	Repos		int		`json:"public_repos"`
	Email		string		`json:"email"`
}

func getUsers(name string) User {
	resp, err := http.Get(apiURL + userEndpoint + name)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}
	if resp.StatusCode != 200{
		log.Fatalf("User doesnot exist status: %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var user User
	json.Unmarshal(body, &user)

	return user
}
