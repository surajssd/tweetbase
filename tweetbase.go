package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	allTweetsFile = "./alltweets.json"
)

func getMaxID(client *twitter.Client) int64 {
	t := &twitter.UserTimelineParams{
		ScreenName: "surajd_",
		Count:      1,
	}
	tweets, _, err := client.Timelines.UserTimeline(t)
	if err != nil {
		log.Fatalf("Error downloading tweets, %v", err)
	}
	log.Println("Got max id:", tweets[0].ID)
	return tweets[0].ID
}

func downloadAllTweets(client *twitter.Client, maxID int64) []twitter.Tweet {
	var allTweets []twitter.Tweet
	var firstID int64

	log.Println("Tweets download started")
	for firstID != maxID {

		t := &twitter.UserTimelineParams{
			ScreenName: "surajd_",
			MaxID:      maxID,
			Count:      2000,
		}

		tweets, _, err := client.Timelines.UserTimeline(t)
		if err != nil {
			log.Fatalf("Error downloading tweets, %v", err)
		}
		allTweets = append(allTweets, tweets...)

		maxID = tweets[len(tweets)-1].ID
		firstID = tweets[0].ID
	}
	log.Println("Tweets download complete")

	return allTweets
}

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	log.Printf("All clients ready")

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		log.Fatalf("Error verifying the user, %v", err)
	}
	log.Println("Credentials verified")

	var tweets []twitter.Tweet
	maxID := getMaxID(client)
	tweets = downloadAllTweets(client, maxID)

	data, err := json.MarshalIndent(tweets, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling all tweets, %v", err)
	}
	err = ioutil.WriteFile(allTweetsFile, data, 0644)
	if err != nil {
		log.Fatalf("Error dumping all tweets to file, %v", err)
	}
	log.Println("Tweets dumped to", allTweetsFile)

	for _, t := range tweets {
		fmt.Printf("\n----------------------------\n%s\n", t.Text)
	}

}
