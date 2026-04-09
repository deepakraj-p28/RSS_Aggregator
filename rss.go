package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		PubDate     string    `xml:"pubDate"`
		Link        string    `xml:"link"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Link        string `xml:"link"`
	Content     string `xml:"content:encoded"`
}

func urlToFeed(url string) (RSSFeed, error) {
	rssFeed := RSSFeed{}

	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return rssFeed, err
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return rssFeed, err
	}

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return rssFeed, err
	}

	return rssFeed, nil
}
