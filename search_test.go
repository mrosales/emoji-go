package emoji

import (
	"reflect"
	"testing"
)

func exactMatch(s string) Info {
	for _, info := range All {
		if info.Name == s || info.Unified == s || info.Character == s {
			return info
		}
	}
	return Info{}
}

func exactMatches(terms ...string) []Info {
	var results []Info
	for _, term := range terms {
		results = append(results, exactMatch(term))
	}
	return results
}

func TestSearch(t *testing.T) {
	type args struct {
		query   string
		options []SearchOption
	}
	tests := []struct {
		name string
		args args
		want []Info
	}{
		{
			"rocket match",
			args{"rocket", []SearchOption{WithLimit(1)}},
			exactMatches("rocket"),
		},
		{
			"rock matches",
			args{"rock", []SearchOption{WithLimit(5), WithMaxDistance(10)}},
			exactMatches("rock", "rocket", "shamrock", "alarm_clock", "timer_clock"),
		},
		{
			"no results",
			args{"fubar", []SearchOption{WithLimit(5), WithMaxDistance(10)}},
			[]Info{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			searcher := NewSearchIndex(tt.args.options...)
			if got := searcher.Search(tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
