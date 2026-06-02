package main

import (
	"context"
	"encoding/json"
	"testing"
)

func TestFetchFeedShape(t *testing.T) {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		t.Fatal(err)
	}

	readable, err := json.MarshalIndent(feed, "", "")
	t.Log(string(readable))
}
