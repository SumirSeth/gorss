package main

import (
	"fmt"

	"github.com/sumirseth/gorss/internal/feeds" 
)

func main() {
	parser := feeds.NewFeedParser()
	feed, items, err := parser.Parse("https://antfu.me/feed.xml")
	if err != nil {
		panic(err)
	}
	fmt.Println("Feed:", feed.Title)
	for _, item := range items {
		fmt.Println("-", item.Title)
	}
}
