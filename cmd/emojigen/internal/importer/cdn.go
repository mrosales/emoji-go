package importer

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

// CDN provides an interface for a remote emoji dataset.
type CDN struct {
	url        url.URL
	httpClient *http.Client
}

// NewCDN creates a new CDN client for emoji data.
func NewCDN() (*CDN, error) {
	u, err := url.Parse("https://cdn.jsdelivr.net/npm/emoji-datasource-apple@14.0.0")
	if err != nil {
		return nil, fmt.Errorf("invalid CDN url: %w", err)
	}
	return &CDN{
		url:        *u,
		httpClient: http.DefaultClient,
	}, nil
}

func (c *CDN) get(ctx context.Context, pathComponents ...string) (*http.Response, error) {
	u := c.url
	u.Path = path.Join(append([]string{u.Path}, pathComponents...)...)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed creating http GET request: %w", err)
	}
	return c.httpClient.Do(req)
}

// DownloadEmojiInfo retrieves and decodes a list of emoji metadata from the CDN.
func (c *CDN) DownloadEmojiInfo(ctx context.Context) ([]EmojiInfo, error) {
	resp, err := c.get(ctx, "emoji.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed fetching emoji metadata with status code %d", resp.StatusCode)
	}
	return ParseEmojiData(resp.Body)
}

// DownloadSpriteSheet downloads the sprite sheet from the CDN.
func (c *CDN) DownloadSpriteSheet(ctx context.Context, width int) (*SpriteSheet, error) {
	switch width {
	case 16, 13, 64:
		break
	default:
		width = 64
	}

	sheet := fmt.Sprintf("/img/apple/sheets-clean/%d.png", width)
	resp, err := c.get(ctx, sheet)
	if err != nil {
		return nil, fmt.Errorf("failed loading sprite data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed fetching emoji metadata with status code %d", resp.StatusCode)
	}
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return nil, fmt.Errorf("failed reading sprite sheet response %w", err)
	}
	return NewSpriteSheet(buf, width)
}
