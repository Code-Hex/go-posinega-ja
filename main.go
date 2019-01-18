package main

import (
	"github.com/ikawaha/kagome/tokenizer"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

}

func Decision(str string) float64 {
	str = Normalize(Clean(str))
	t := tokenizer.New()
	tokens := t.Tokenize(str)
	words := make([]string, 0, len(tokens))
	for _, token := range tokens {
		word := token.Surface
		if IsStopWord(word) {
			continue
		}
		if token.Class != tokenizer.DUMMY {
			words = append(words, word)
		}
	}
	return GetScore(words...)
}
