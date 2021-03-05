package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s output_file", os.Args[0])
	}
	outputFile := os.Args[1]
	emojis, err := DownloadEmojiInfo()
	if err != nil {
		log.Fatalf("failed parsing emoji info: %v", err)
	}
	log.Printf("successfully downloaded %d emojis", len(emojis))

	buf := &bytes.Buffer{}
	if err := keywordTemplate.Execute(buf, map[string]interface{}{
		"Package": "emoji",
		"Emojis":  emojis,
	}); err != nil {
		log.Fatalf("failed rendering template: %v", err)
	}

	dirname := path.Dir(outputFile)
	if err := os.MkdirAll(dirname, 0777); err != nil {
		log.Fatalf("failed creating output directory %s: %v", dirname, err)
	}
	if err := ioutil.WriteFile(outputFile, buf.Bytes(), 0777); err != nil {
		log.Fatalf("failed writing file: %v", err)
	}
	log.Printf("completed writing data to %s", outputFile)
}
