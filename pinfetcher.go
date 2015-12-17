package main

import "flag"
import "time"
import "log"
import "net/http"
import "encoding/json"
import "text/template"
import "os"
import "regexp"
import "strings"
import "net/url"

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
	TagArray []string
}

type Pin struct {
	Href string
	Description string
	Extended string
	Tags []string
}

func main() {

	// set-up command line flags
	apiKeyPtr := flag.String("api-key", "", "your PinBoard API key")
	daysOffsetPtr := flag.Int("d", 7, "number of days to retrieve")
	templateFilePtr := flag.String("t", "default.tpl", "template file")
	tagsPtr := flag.String("tags", "", "a space separated list of up to three tags")
	flag.Parse()


	// confirm we have a valid API key
	r, _ := regexp.Compile("^[[:alnum:]]*:[0-9A-F]*")
	if !r.MatchString(*apiKeyPtr) {
		log.Fatal("Invalid API key provided.")
	}


	// calculate the the date range to fetch
	currentTime := time.Now().Local()
	toTime := currentTime.AddDate(0, 0, -1)
	fromTime := currentTime.AddDate(0, 0, -1 * *daysOffsetPtr -1)


	// round timestamps to nearest 24 hours and convert to strings
	toTimeString := toTime.Round(time.Hour*24).Format(time.RFC3339)
	fromTimeString := fromTime.Round(time.Hour*24).Format(time.RFC3339)


	// construct the url
	u, err := url.Parse("https://api.pinboard.in/v1/posts/all")
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Add("auth_token", *apiKeyPtr)
	q.Add("fromdt", fromTimeString)
	q.Add("todt", toTimeString)
	q.Add("format", "json")

	// format the tags (if provided) for passing to the API
	if *tagsPtr != "" {
		tags := prepareTags(tagsPtr)
		q.Add("tag", tags)
	}

	u.RawQuery = q.Encode()


	// fetch latest pins
	data := []PinJson{}
    err = getJson(u, &data)
    if err != nil {
    	log.Fatal(err)
    }


    // structure the pins
    for i := range data {
  		pin := &data[i] // we need to reference the original slice
  		pin.TagArray = strings.Split(pin.Tags, " ")
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
func getJson(u *url.URL, target interface{}) error {
    r, err := http.Get(u.String())
    if err != nil {
        return err
    }

    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}


func prepareTags(tagsPtr *string) string {
	tags := strings.Split(*tagsPtr, " ")
	if len(tags) > 3 {
		tags = tags [:3]
		log.Print("The Pinboard API supports a maximum of 3 tags. Only the first three will be used: ", tags)
	}
	return strings.Join(tags, " ")
}
