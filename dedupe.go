package dedupe

import (
	"slices"
	"sort"
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
	for _, word := range words {
		d.AddWord(word)
	}
}

func (d *Data) AddWord(word string) {
	d.words = append(d.words, strings.ToLower(word))
}

func (d *Data) Deduped() []string {
	return d.WordsDeduped
}

// func (d *Data) Contains(word []string) {
// 	d.words = slices.ContainsFunc(d.words, word...)
// }

// func (d *Data) DedupeWords() {
// 	// for _, word := range d.words {
// 	// 	d.Dedupe1(word)
// 	// }
// 	d.Dedupe2()
// }

func (d *Data) Dedupe1() {
	data := d.words[:0]
	for _, word := range d.words {
		if !slices.ContainsFunc(d.words, func(e string) bool { return e == word }) {
			d.Count[word] += 1
			data = append(data, word)
		}
	}

	d.WordsDeduped = data
}

// func (d *Data) Sort() {
// 	// slices.Sort(d.words)
// 	sort.Slice(d.words, func(i, j string) bool {
// 		return a[i] < a[j]
// 	})
// }

// Dedupe - dedupe array of strings
func (d *Data) Dedupe2() {
	set := make(map[string]interface{})
	var res []string
	sort.Slice(d.words, func(i, j int) bool { return d.words[i] > d.words[j] })
	// sort.Sort(sort.Reverse(sort.StringSlice{}d.words))

	for _, sub := range d.words {
		// strs := strings.Split(s, " ")
		// for _, sub := range s {
		if _, ok := set[sub]; !ok {
			set[sub] = nil
			res = append(res, sub)
		}
		// }
	}

	d.WordsDeduped = res
}
