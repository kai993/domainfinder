package thesaurus

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BighugelabsURL :
const BighugelabsURL = "http://words.bighugelabs.com/api/2/%s/%s/json"

type (
	BigHuge struct {
		APIKey string
	}

	synonyms struct {
		Noun *words `json:"noun"`
		Verb *words `json:"verb"`
	}

	words struct {
		Syn []string `json:"syn"`
	}
)

// Synonyms : 同義語を検索する
func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	url := fmt.Sprintf(BighugelabsURL, b.APIKey, term)
	response, err := http.Get(url)
	if err != nil {
		return syns, fmt.Errorf("bighuge: %qの類語検索に失敗しました: %v", term, err)
	}

	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}

	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)

	return syns, nil
}
