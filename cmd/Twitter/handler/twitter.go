package handler

import (
	"CliTaskManager/cmd/Twitter/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTweets(userName string) {
	var tweets model.SearchTweetsModel
	url := fmt.Sprintf("https://api.twitter.com/2/tweets/search/recent?query=from:%s", userName)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer AAAAAAAAAAAAAAAAAAAAAEA6bQEAAAAAyCDWnB1jI8gdBFazOzNd%2FhFh%2F3I%3DWcYvAa60XuBoIOit53stsc1agp6SKdg2shOE1RObOYY7wCjnhP"},
	}
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &tweets)

	if err != nil {
		fmt.Println(err)
	}
	for _, v := range tweets.Data {
		fmt.Printf("id: %s, tweet: %s\n", v.ID, v.Text)
	}
}

func GetFollowingAccounts(userId string) {
	var followingAccount model.FollowingAccounts
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/following", userId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer AAAAAAAAAAAAAAAAAAAAAEA6bQEAAAAAyCDWnB1jI8gdBFazOzNd%2FhFh%2F3I%3DWcYvAa60XuBoIOit53stsc1agp6SKdg2shOE1RObOYY7wCjnhP"},
	}
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &followingAccount)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range followingAccount.Data {
		fmt.Printf("id: %s, name: %s, username: %s\n", v.ID, v.Name, v.UserName)
	}
}

func GetFollowers(userId string) {
	var followingAccount model.FollowingAccounts
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/followers", userId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer AAAAAAAAAAAAAAAAAAAAAEA6bQEAAAAAyCDWnB1jI8gdBFazOzNd%2FhFh%2F3I%3DWcYvAa60XuBoIOit53stsc1agp6SKdg2shOE1RObOYY7wCjnhP"},
	}
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &followingAccount)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range followingAccount.Data {
		fmt.Printf("id: %s, name: %s, username: %s\n", v.ID, v.Name, v.UserName)
	}
}
