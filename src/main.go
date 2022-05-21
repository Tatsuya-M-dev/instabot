package main

import (
	"fmt"

	"example.com/instabot/api"
)

func main() {
	fmt.Println("Hello world!")
	api.Tweet("杏樹……")
	// if err := scrape(); err != nil {
	// 	fmt.Printf("error %v", err)
	// }

}

func scrape() error {
	accounts, err := getInstagramAccounts()
	if err != nil {
		return fmt.Errorf("failed to get instagram accounts")
	}

	for _, accountId := range accounts {
		err := notifyTweet(accountId)
		if err != nil {
			return fmt.Errorf("failed to notify tweet %v", err)
		}
	}
	return nil
}

func getInstagramAccounts() ([]string, error) {
	ret := []string{}
	return ret, nil
}

func notifyTweet(instagramId string) error {
	postIdArray, err := api.CheckInstagram(instagramId)
	if err != nil {
		return fmt.Errorf("failed to check instagram: %v", err)
	}
	for _, postId := range postIdArray {
		data, err := api.GetPostData(postId)
		if err != nil {
			return fmt.Errorf("failed to get post data %v", err)
		}
		if err := api.Tweet(data.Text); err != nil {
			return fmt.Errorf("failed to tweet %v", err)
		}
		if err := api.SavePostData(data); err != nil {
			return fmt.Errorf("failed to save post data")
		}
	}
	fmt.Printf("no implemention")
	return nil
}
