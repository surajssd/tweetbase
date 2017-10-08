package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	consumerKey    = "2d27tn94qw0e7f3w25VMS84Gc"
	consumerSecret = "tLu3tvKMjvYMq9EyrOiOyBscJcBI8JxV1yuDIinNHo8SxS2aA0"
	accessToken    = "285370024-vkYNqKeqtpT6D6fZsNsyNyJNvwWM9sC4uW0ljEKe"
	accessSecret   = "5n9WyLyZgHVXhqU8CHjoZBVSc12Aj5ZSpYhmLyTiOKinR"
)

func main() {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	//tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
	//	Count: 20,
	//})
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	// Send a Tweet
	// tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)

	// Status Show
	tweet, resp, err := client.Statuses.Show(585613041028431872, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Tweets: ", tweet)
	fmt.Println("response: %v", resp)

	// Search Tweets
	//search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
	//	Query: "gopher",
	//})

	// User Show
	//user, resp, err := client.Users.Show(&twitter.UserShowParams{
	//	ScreenName: "dghubble",
	//})

	// Followers
	//followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})
}
