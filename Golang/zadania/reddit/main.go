package main

import (
	"io"
	"log"
	"os"
	"reddit/fetcher"
	"time"
)

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change

	f = fetcher.NewRedditClient("https://www.reddit.com/r/golang.json", 5*time.Second)
	w, _ = os.Create("redditoutput.txt")

	err := f.Fetch()
	if err != nil {
		log.Fatalf("error occured during fetching: %v", err)
	}
	err = f.Save(w)
	if err != nil {
		log.Fatalf("error occured during saving: %v", err)
	}
}
