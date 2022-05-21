package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	if err := scrape(); err != nil {
		fmt.Printf("error %v", err)
	}

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
	fmt.Printf("no implemention")
	return nil
}
