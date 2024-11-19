package dedupe

import (
	"fmt"
	"slices"
	"strings"
)

type Data struct {
	Name         string
	words        []string
	WordsDeduped []string

	Count map[string]int
}

func NewSet() *Data {
	return &Data{
		Count: make(map[string]int),
	}
}

func (d *Data) AddWords(words []string) {
	d.words = append(d.words, words...)
}

func (d *Data) AddWord(word string) {
	d.words = append(d.words, word)
	data := d.words[:0]
	if !slices.ContainsFunc(data, func(e string) bool { return e == word }) {
		d.Count[word] += 1
		data = append(data, word)
	}
	fmt.Println("data", data)
}

func (d *Data) Deduped() []string {
	return d.WordsDeduped
}

// func (d *Data) Contains(word []string) {
// 	d.words = slices.ContainsFunc(d.words, word...)
// }

func (d *Data) DedupeWords() {
	// for _, word := range d.words {
	// 	d.Dedupe1(word)
	// }
	d.Dedupe2(d.words)
}

func (d *Data) Dedupe1(word string) {
	data := d.words[:0]
	if !slices.ContainsFunc(d.words, func(e string) bool { return e == word }) {
		d.Count[word] += 1
		data = append(data, word)
	}

	d.WordsDeduped = data
}

// Dedupe - dedupe array of strings
func (d *Data) Dedupe2(words []string) {
	//  dedupe????
	set := make(map[string]interface{})
	var res []string

	for _, s := range words {
		strs := strings.Split(s, " ")
		for _, sub := range strs {
			st := strings.ToLower(sub)
			if _, ok := set[st]; !ok {
				set[st] = nil
				res = append(res, st)
			}
		}
	}

	d.WordsDeduped = res
}
