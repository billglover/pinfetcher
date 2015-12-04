package main

import "flag"
import "fmt"

func main() {
	apiKeyPtr := flag.String("api-key", "", "your PinBoard API key")

	flag.Parse()

	fmt.Println("api-key:", *apiKeyPtr)
}	