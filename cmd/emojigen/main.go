// Package main implements a CLI that downloads emoji data and images.
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mrosales/emoji-go/cmd/emojigen/internal/importer"
)

func main() {
	var (
		timeout       time.Duration
		datasetOutput string
		imageOutput   string
	)
	flag.DurationVar(
		&timeout,
		"timeout",
		30*time.Second,
		"timeout for the network requests")
	flag.StringVar(
		&datasetOutput,
		"dataset",
		"",
		"file to write generated data to")
	flag.StringVar(
		&imageOutput,
		"images",
		"",
		"directory to write emoji images to")

	flag.Parse()

	cdn, err := importer.NewCDN()
	if err != nil {
		log.Fatalf("failed creating cdn client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	emojis, err := cdn.DownloadEmojiInfo(ctx)
	if err != nil {
		log.Fatalf("failed parsing emoji info: %v", err)
	}
	log.Printf("successfully downloaded %d emojis", len(emojis))

	if len(datasetOutput) > 0 {
		if err := writeDataset(datasetOutput, emojis); err != nil {
			log.Fatalf("failed writing dataset: %v", err)
		}
		log.Printf("successfully wrote emoji dataset to %s", datasetOutput)
	}

	if len(imageOutput) > 0 {
		sprites, err := cdn.DownloadSpriteSheet(ctx, 64)
		if err != nil {
			log.Fatalf("failed downloading sprites: %v", err)
		}
		n, err := writeSprites(imageOutput, emojis, sprites)
		if err != nil {
			log.Fatalf("failed writing images: %v", err)
		}
		log.Printf("successfully wrote %d emoji images to %s", n, imageOutput)
	}

}

func writeDataset(output string, emojis []importer.EmojiInfo) error {
	buf := &bytes.Buffer{}
	if err := importer.RenderTemplate(buf, "emoji", emojis); err != nil {
		return fmt.Errorf("failed rendering template: %w", err)
	}
	dirname := filepath.Dir(output)
	if err := os.MkdirAll(dirname, 0777); err != nil {
		return fmt.Errorf("failed creating output directory %s: %v", dirname, err)
	}
	if err := ioutil.WriteFile(output, buf.Bytes(), 0777); err != nil {
		return fmt.Errorf("failed writing file: %v", err)
	}
	log.Printf("completed writing data to %s", output)
	return nil
}

func writeSprites(output string, emojis []importer.EmojiInfo, sprites *importer.SpriteSheet) (int, error) {
	if err := os.MkdirAll(output, 0777); err != nil {
		return 0, fmt.Errorf("failed creating output directory %s: %v", output, err)
	}
	count := 0
	for _, info := range emojis {
		for _, imageData := range allImages(info) {
			imageOutPath := filepath.Join(output, imageData.Image)
			image := sprites.Get(imageData.SheetX, imageData.SheetY)
			f, err := os.Create(imageOutPath)
			if err != nil {
				return count, fmt.Errorf("failed creating output file %s: %w", imageOutPath, err)
			}
			if err := png.Encode(f, image); err != nil {
				_ = f.Close()
				return count, fmt.Errorf("failed encoding PNG %s: %w", imageOutPath, err)
			}
			if err := f.Close(); err != nil {
				return count, fmt.Errorf("failed closing file %s: %w", imageOutPath, err)
			}
			count++
		}
	}
	return count, nil
}

func allImages(info importer.EmojiInfo) []importer.EmojiImageData {
	output := []importer.EmojiImageData{info.EmojiImageData}
	for _, variation := range info.SkinVariations {
		output = append(output, variation)
	}
	return output
}
