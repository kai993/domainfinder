package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kai993/domainfinder/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	var thesaurus thesaurus.Thesaurus = &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalf("%qの類語検索に失敗しました: %v\n", word, err)
		}

		if len(syns) == 0 {
			log.Fatalf("%qには類語はありませんでした\n", word)
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}