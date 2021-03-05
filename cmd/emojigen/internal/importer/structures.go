package importer

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

// ParseEmojiData parses emoji information from a JSON reader.
func ParseEmojiData(r io.Reader) ([]EmojiInfo, error) {
	var emojis []EmojiInfo
	if err := json.NewDecoder(r).Decode(&emojis); err != nil {
		return nil, err
	}

	var output []EmojiInfo
	for _, info := range emojis {
		info.Name = sanitizedName(info)
		info.ShortName = sanitizedShortName(info)
		info.Unified = strings.ToLower(info.Unified)

		chr, err := loadUnifiedSequence(info.Unified)
		if err != nil {
			return nil, fmt.Errorf("invalid emoji sequence: \"%s\": %w", info.Unified, err)
		}
		info.Character = chr

		var alternateNames []string
		hasName := map[string]bool{}
		for _, keyword := range info.ShortNames {
			if hasName[keyword] {
				continue
			}
			alternateNames = append(alternateNames, keyword)
			hasName[keyword] = true
		}
		info.ShortNames = alternateNames

		mutatedVariations := map[string]EmojiImageData{}
		for modifier, variation := range info.SkinVariations {
			switch modifier {
			case "1F3FB", "1F3FC", "1F3FD", "1F3FE", "1F3FF":
				break
			default:
				// TODO support multi-skin-tone emojis
				continue
			}
			variationChr, err := loadUnifiedSequence(variation.Unified)
			if err != nil {
				return nil, fmt.Errorf("invalid emoji sequence: \"%s\": %w", variation.Unified, err)
			}
			variation.Character = variationChr
			mutatedVariations[modifier] = variation
		}
		info.SkinVariations = mutatedVariations
		output = append(output, info)
	}
	return output, nil
}

// EmojiInfo captures standard emoji attributes for a named and tagged emoji.
type EmojiInfo struct {
	Name       string   `json:"name"`
	ShortName  string   `json:"short_name"`
	ShortNames []string `json:"short_names"`
	Text       string   `json:"text"`
	Texts      []string `json:"texts"`
	Category   string   `json:"category"`
	SortOrder  int      `json:"sort_order"`
	EmojiImageData
	SkinVariations map[string]EmojiImageData `json:"skin_variations"`
}

// EmojiImageData captures standard emoji attributes.
// These details are shared by both top-level emojis and variations.
type EmojiImageData struct {
	// A hyphen separated sequence of hex-encoded codepoints.
	// Includes zero-width-joiner character for multi code sequences.
	Unified string `json:"unified"`
	// Character is the actual emoji character consisting of one or more codepoints.
	// Not present in original input but will be populated by parser
	Character string `json:"character,omitempty"`
	// NonQualified is set if the emoji also has usage without a variation selector.
	NonQualified string `json:"non_qualified"`
	// Image is the image name as available via the CDN.
	Image string `json:"image"`
	// SheetX is the X index of the image in the sprite sheet.
	SheetX int `json:"sheet_x"`
	// SheetY is the Y index of the image in the sprite sheet.
	SheetY int `json:"sheet_y"`
	// Ignore legacy codepoint data
	// Docomo         string `json:"docomo"`
	// AU             string `json:"au"`
	// Softbank       string `json:"softbank"`
	// Google         string `json:"google"`
	AddedIn        string `json:"added_in"`
	HasImgApple    bool   `json:"has_img_apple"`
	HasImgGoogle   bool   `json:"has_img_google"`
	HasImgTwitter  bool   `json:"has_img_twitter"`
	HasImgFacebook bool   `json:"has_img_facebook"`
	Obsoletes      string `json:"obsoletes"`
	ObsoletedBy    string `json:"obsoleted_by"`
}

// loadUnifiedSequence returns an emoji unicode string from a unified hex sequence.
//
// A sequence is hyphen separated sequence of hex-encoded UTF8 codepoints.
// As an example, "2708-fe0f" represents ✈️
func loadUnifiedSequence(unified string) (string, error) {
	hexChars := strings.Split(unified, "-")
	output := make([]byte, 0, len(hexChars)+len(hexChars)-1)
	buf := make([]byte, 4)

	for _, hexChar := range hexChars {
		if len(hexChar) == 0 {
			return "", fmt.Errorf("invalid hex sequence %s", unified)
		}
		intVal, err := strconv.ParseUint("0x"+hexChar, 0, 64)
		if err != nil {
			return "", err
		}

		runeVal := rune(intVal)

		if !utf8.ValidRune(runeVal) {
			return "", fmt.Errorf("invalid utf8 rune from \"%s\"", unified)
		}
		output = append(output, buf[0:utf8.EncodeRune(buf, runeVal)]...)
	}
	return string(output), nil
}
