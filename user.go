package main

import (
	"encoding/json"
	"time"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiURl= "https://api.gitub.com"
	userEndpoint = "/users/"
)

type User struct {
	Login		string		`json:"login"`
	ID		int		`json:"id"`
}

func getUsers(name string) User {
	resp, err := http.Get(apiURL + userEndpoint + name)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer.resp.Body.close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %ss\n", err)
	}

	var user User
	json.Umarshal(body, &user)

	return user
