package dedupe

import (
	"fmt"
	"slices"
)

type Data struct {
	Name         string
	Words        []string
	WordsDeduped []string

	Count map[string]int
}

func NewSet() *Data {
	return &Data{
		Count: make(map[string]int),
	}
}

func (d *Data) AddWords(words []string) {
	d.Words = append(d.Words, words...)
}

func (d *Data) AddWord(word string) {
	d.Words = append(d.Words, word)
	data := d.Words[:0]
	if !slices.ContainsFunc(data, func(e string) bool { return e == word }) {
		d.Count[word] += 1
		data = append(data, word)
	}
	fmt.Println("data", data)
}

func (d *Data) Deduped() []string {
	return d.WordsDeduped
}

func (d *Data) Contains(word []string) {
	d.Words = append(d.Words, word...)
}

func (d *Data) Dedupe(words []string) {
	data := d.Words[:0]
	for _, word := range words {
		if !slices.ContainsFunc(data, func(e string) bool { return e == word }) {
			d.Count[word] += 1
			data = append(data, word)
		}
	}
	d.WordsDeduped = data
}
