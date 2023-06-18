package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch() error
	Save(io.Writer) error
}

type RedditClient struct {
	responseData response
	host         string
	c            *http.Client
}

func NewRedditClient(host string, timeout time.Duration) *RedditClient {
	return &RedditClient{
		host:         host,
		responseData: response{},
		c: &http.Client{
			Timeout: timeout,
		},
	}
}

func (rc *RedditClient) Fetch() error {
	req, err := http.NewRequest(http.MethodGet, rc.host, nil)
	if err != nil {
		return fmt.Errorf("error occured while creating request %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := rc.c.Do(req)
	if err != nil {
		return fmt.Errorf("error occured while getting data %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&rc.responseData)
	if err != nil {
		return fmt.Errorf("error occured while decoding response body: %v", err)
	}

	log.Printf("successfully fetched data from %s\n", rc.host)

	return nil
}

func (rc *RedditClient) Save(w io.Writer) error {
	if len(rc.responseData.Data.Children) == 0 {
		return fmt.Errorf("response data contained no children elements")
	}

	for _, v := range rc.responseData.Data.Children {
		_, err := fmt.Fprintf(w, "Title: %s\nURL: %s\n", v.Data.Title, v.Data.URL)
		if err != nil {
			return fmt.Errorf("error occured while writing data: %v", err)
		}
	}

	return nil
}
