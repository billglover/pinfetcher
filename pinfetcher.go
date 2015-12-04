package main

import "flag"
import "fmt"
import "time"

func main() {
	// --api-key
	apiKeyPtr := flag.String("api-key", "", "your PinBoard API key")
	daysOffsetPtr := flag.Int("d", 7, "number of days to retrieve")
	flag.Parse()

	// time
	currentTime := time.Now().Local()
	yesterdayTime := currentTime.AddDate(0, 0, -1)
	lastWeekTime := currentTime.AddDate(0, 0, -1 * *daysOffsetPtr -1)
	yesterday := yesterdayTime.Round(time.Hour*24).Format(time.RFC3339)
	lastWeek := lastWeekTime.Round(time.Hour*24).Format(time.RFC3339)
	fmt.Println("From: ", lastWeek)
	fmt.Println("To: ", yesterday)

	// url
	url := fmt.Sprintf("https://api.pinboard.in/v1/posts/all?auth_token=%s&fromdt=%s&todt=%s&format=json", *apiKeyPtr, lastWeek, yesterday)
	fmt.Println(url)

}	