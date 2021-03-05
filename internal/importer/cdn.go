package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	emojiCDNURL = "https://cdn.jsdelivr.net/npm/emoji-datasource-apple@6.0.0"
)

// DownloadEmojiInfo retrieves and decodes a list of emoji metadata from the CDN.
func DownloadEmojiInfo() ([]EmojiInfo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/emoji.json", emojiCDNURL))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed fetching emoji metadata with status code %d", resp.StatusCode)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %w", err)
	}
	return ParseEmojiData(data)
}
