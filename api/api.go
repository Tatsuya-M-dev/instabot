package api

import (
	"fmt"
	"time"
)

type TweetData struct {
	text string
}

func Tweet(text string) error {
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
