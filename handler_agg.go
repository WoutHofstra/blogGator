package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/xml"
	"context"
	"html"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}


func handlerAgg(s *state, cmd command) error {

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		fmt.Println(err)
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Println("Collecting feeds every", timeBetweenRequests)
	for ;; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil
}


func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error: couldnt get request: %w", err)
	}

	req.Header.Set("User-Agent", "blogGator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error: couldnt get response: %w", err)
	}
        defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error: couldnt read response body: %w", err)
	}


	var feed RSSFeed


	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, fmt.Errorf("Error: Couldnt unmarshal data: %w", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, _ := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
                feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}


	return &feed, nil

}

func scrapeFeeds(s *state) {

	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		fmt.Println(err)
	}
	s.db.MarkFeedFetched(ctx, feed.ID)
	content, err := fetchFeed(ctx, feed.Url)
	if err != nil {
                fmt.Println(err)
        }
	for _, c := range content.Channel.Item {
		fmt.Println(c.Title)
	}
}
