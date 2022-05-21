package api

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TweetData struct {
	text string
}

func Tweet(text string) error {
	config := oauth1.NewConfig(consumer_key, consumer_key_secret)
	token := oauth1.NewToken(access_token, access_token_secret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Send a Tweet
	_, _, err := client.Statuses.Update(text, nil)
	if err != nil {
		return fmt.Errorf("failed to tweet %v", err)
	}
	return nil
}

//get new post id
func CheckInstagram(instagramId string) ([]string, error) {
	// publish instagram API

	// if no new post
	if false {
		return nil, fmt.Errorf("no new post")
	}
	ret := []string{""}
	return ret, nil
}

type InstagramPostData struct {
	PostId   string
	Text     string
	CreateAt time.Time
}

func GetPostData(postId string) (*InstagramPostData, error) {
	return &InstagramPostData{}, nil
}

func SavePostData(data *InstagramPostData) error {
	return nil
}

func HttpPost(url, token, text string) error {
	jsonStr := `{"token":"` + token + `","text":"` + text + `"}`

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
