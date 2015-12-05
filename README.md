# pinfetcher
Turn Pinboard links into Markdown based on a simple template.

## Usage

```
go run pinfetcher.go -h
Usage of pinfetcher:
  -api-key string
    	your PinBoard API key
  -d int
    	number of days to retrieve (default 7)
  -t string
    	template file (default "default.tpl")
```

## API Key

A Pinboard authentication token is a short opaque identifier in the form "username:TOKEN".

Users can find their API token on their [settings](https://pinboard.in/settings/password) page.
