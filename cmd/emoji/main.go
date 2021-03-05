package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/mrosales/emoji-go"
	"github.com/spf13/cobra"
)

func main() {
	var (
		searchOptLimit       = 0
		searchOptMaxDistance = 10
		outputFormat         = "char"
		skinTone             = ""
		formatters           = map[string]func([]emoji.Info, emoji.Modifier) (string, error){
			"char": charFormatter,
			"json": jsonFormatter,
			"text": textFormatter,
		}
	)

	root := &cobra.Command{
		Use:     "search [-l limit] [-d maxdistance] [-f char|text|json] [-s skin] query",
		Short:   "Look up emojis with a keyword",
		Aliases: []string{"s"},
		Example: "search -l 1 rocket",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.SetOut(os.Stdout)
			query := strings.Join(args, " ")
			searcher := emoji.NewSearchIndex(
				emoji.WithLimit(searchOptLimit),
				emoji.WithMaxDistance(searchOptMaxDistance))

			results := searcher.Search(query)
			formatter, ok := formatters[outputFormat]
			if !ok {
				cmd.PrintErr("unsupported output format \"%s\"", outputFormat)
				_ = cmd.Usage()
				os.Exit(1)
			}

			modifier := emoji.Modifier(0)
			modifier, err := emoji.NewModifier(skinTone)
			if err != nil {
				cmd.PrintErrf("unsupported skin tone %s: %v", skinTone, err)
				os.Exit(1)
			}

			output, err := formatter(results, modifier)
			if err != nil {
				cmd.PrintErrf("failed formatting output: %v", err)
				os.Exit(1)
			}

			cmd.Printf("%s", string(output))
		},
	}

	root.PersistentFlags().IntVarP(
		&searchOptLimit,
		"limit",
		"l",
		searchOptLimit,
		"limit returned entries, or 0 for no limit")
	root.PersistentFlags().IntVarP(
		&searchOptMaxDistance,
		"maxdistance",
		"d",
		searchOptMaxDistance,
		"maximum Levenshtein distance for keyword matches")
	root.PersistentFlags().StringVarP(
		&skinTone,
		"skin",
		"s",
		"",
		"skin tone [light|medium_light|medium|medium_dark|dark]")
	root.PersistentFlags().StringVarP(
		&outputFormat,
		"format",
		"f",
		outputFormat,
		"output format [char|json|text]")

	if err := root.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "failed executing command: %v", err)
		os.Exit(1)
	}
}

func charFormatter(results []emoji.Info, skinTone emoji.Modifier) (string, error) {
	sb := strings.Builder{}
	sb.Grow(4 * len(results))
	for i, info := range results {
		imageData := info.ImageForModifier(skinTone)
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(imageData.Character)
	}
	return sb.String(), nil
}

func jsonFormatter(results []emoji.Info, skinTone emoji.Modifier) (string, error) {
	data, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func textFormatter(results []emoji.Info, skinTone emoji.Modifier) (string, error) {
	buf := &bytes.Buffer{}
	w := tabwriter.NewWriter(buf, 4, 0, 2, ' ', 0)
	for _, info := range results {
		imageData := info.ImageForModifier(skinTone)
		line := strings.Join(
			[]string{
				info.Name,
				imageData.Unified,
				imageData.Character,
			},
			"\t")
		fmt.Fprintln(w, line)
	}
	if err := w.Flush(); err != nil {
		return "", err
	}
	return buf.String(), nil
}
