package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

var user string

func main() {
	fmt.Println("\nWelcome to the CLI\n")
	flag.Parse()
	if flag.NFlag() == 0 {
		printUsage()
	}
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n\n", users)
	i := 1
	for _, u := range users {
		result := getUsers(u)

		if len(users) > 1 && i-1 != len(users) {
			fmt.Println("\n User ", i)
		}

		fmt.Println(`username:	`, result.Login)
		fmt.Println(`id:		`, result.ID)
		fmt.Println(`url:		`, result.Html_Url)
		fmt.Println(`Repos:		`, result.Repos)
		fmt.Println(`Email:		`, result.Email)
		//fmt.Println(`Users:		`,len(users))
		i = i + 1
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func printUsage() {
	fmt.Printf("usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	fmt.Println("\nPlease add -u plus urgument(username being searched)")
	fmt.Println("OR add -user=urgument(username being searched)")
	os.Exit(1)
}
