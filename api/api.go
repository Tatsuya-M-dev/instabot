package api

import (
	"fmt"
	"time"
)

type TweetData struct {
	text string
}

func tweet(text string) error {
	return nil
}

//get new post id
func checkInstagram(instagramId string) ([]string, error) {
	// publish instagram API

	// if no new post
	if false {
		return nil, fmt.Errorf("no new post")
	}
	ret := []string{""}
	return ret, nil
}

type InstagramPostData struct {
	postId   string
	text     string
	createAt time.Time
}

func getPostData(postId string) {

}

func savePostData(data InstagramPostData) error {
	return nil
}
