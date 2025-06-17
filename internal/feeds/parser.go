package feeds

import (
	"time"
	"github.com/mmcdole/gofeed"

	"github.com/sumirseth/gorss/internal/models"
)

type FeedParser struct {
	parser *gofeed.Parser
}

func NewFeedParser() *FeedParser {
	return &FeedParser{
		parser: gofeed.NewParser(),
	}
}

func (fp *FeedParser) Parse(url string) (*models.Feed, []models.Item, error) {
	feed, err := fp.parser.ParseURL(url)
	if err != nil {
		return nil, nil, err
	}

	feedModel := &models.Feed{
		ID:        generateID(feed.Link), // hash or UUID
		URL:       url,
		Title:     feed.Title,
		UpdatedAt: time.Now(),
	}

	var items []models.Item
	for _, i := range feed.Items {
		items = append(items, models.Item{
			ID:        generateID(i.GUID + i.Link), // deduplication hash
			FeedID:    feedModel.ID,
			Title:     i.Title,
			Link:      i.Link,
			Summary:   i.Description,
			PublishedAt: parseTime(i.PublishedParsed),
			Read:      false,
		})
	}

	return feedModel, items, nil
}

func parseTime(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Now()
}
