package main

import "flag"
import "fmt"
import "time"
import "log"
import "net/http"
import "encoding/json"
import "text/template"
import "os"

// pinboard data structures
type PinJson struct {
	Href string
	Description string
	Extended string
	Meta string
	Hash string
	Time string
	Shared string
	ToRead string
	Tags string
}

func main() {

	// set-up command line flags
	apiKeyPtr := flag.String("api-key", "", "your PinBoard API key")
	daysOffsetPtr := flag.Int("d", 7, "number of days to retrieve")
	templateFilePtr := flag.String("t", "default.tpl", "template file")
	flag.Parse()

	// calculate the the date range to fetch
	currentTime := time.Now().Local()
	toTime := currentTime.AddDate(0, 0, -1)
	fromTime := currentTime.AddDate(0, 0, -1 * *daysOffsetPtr -1)

	// round timestamps to nearest 24 hours and convert to strings
	toTimeString := toTime.Round(time.Hour*24).Format(time.RFC3339)
	fromTimeString := fromTime.Round(time.Hour*24).Format(time.RFC3339)

	// construct the url
	url := fmt.Sprintf("https://api.pinboard.in/v1/posts/all?auth_token=%s&fromdt=%s&todt=%s&format=json", *apiKeyPtr, fromTimeString, toTimeString)

	// fetch latest pins
	data := []PinJson{}
    err := getJson(url, &data)
    if err != nil {
    	log.Fatal(err)
    }

    // print markdown
    t := template.New(*templateFilePtr)
    t, _ = t.ParseFiles(*templateFilePtr)
    err = t.Execute(os.Stdout, data)
    if err != nil {
        log.Fatal(err)
    }
}

// source: http://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang
func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }

    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}
