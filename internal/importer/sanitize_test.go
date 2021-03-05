package main

import (
	"reflect"
	"testing"
)

func Test_uniqueKeywords(t *testing.T) {
	tests := []struct {
		name  string
		input EmojiInfo
		want  []string
	}{
		{
			name: "no name",
			input: EmojiInfo{
				Name:      "emoji name",
				ShortName: "emoji_name",
				ShortNames: []string{
					"emoji_name",
					"something_else",
					"its_an_emoji",
					"<3",
				},
				Text:  "<3",
				Texts: []string{"<3"},
			},
			want: []string{
				"emoji name",
				"emoji_name",
				"something_else",
				"its_an_emoji",
				"<3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueKeywords(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniqueKeywords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sanitizedName(t *testing.T) {
	tests := []struct {
		name  string
		input EmojiInfo
		want  string
	}{
		{
			"should lowercase",
			EmojiInfo{Name: "UPPER CASE NAME"},
			"upper case name",
		},
		{
			"replace underscores",
			EmojiInfo{Name: "short_name"},
			"short name",
		},
		{
			"no duplicate whitespace",
			EmojiInfo{Name: "short__name"},
			"short name",
		},
		{
			"shortname fallback",
			EmojiInfo{ShortName: "short_name"},
			"short name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizedName(tt.input); got != tt.want {
				t.Errorf("sanitizedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
