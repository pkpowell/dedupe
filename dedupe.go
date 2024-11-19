package dedupe

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"sync"
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

type Node struct {
	Children Children      `json:"children,omitempty"`
	Count    int           `json:"count,omitempty"`
	mtx      *sync.RWMutex `json:"-"`
}

type Children map[string]*Node

// newTrie initializes a new Trie
func newTrie() *Node {
	return &Node{
		Children: make(Children),
		// IsEnd:       false,
		Count: 0,
		mtx:   new(sync.RWMutex),
	}
}

// Update updates a word in the trie
func (root *Node) update(word string) {
	root.mtx.Lock()
	defer root.mtx.Unlock()

	current := root

	for _, letter := range word {
		_, ok := current.Children[string(letter)]
		if !ok {
			current.Children[string(letter)] = newTrie()
		}
		current = current.Children[string(letter)]
	}
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
	set := make(map[string]struct{})
	var res = d.words[:0]
	// sort.Slice(d.words, func(i, j int) bool { return d.words[i] > d.words[j] })
	// sort.Sort(sort.Reverse(sort.StringSlice(d.words)))
	slices.SortFunc(d.words, func(i, j string) int {
		return cmp.Compare(len(j), len(i))
	})

	for _, word := range d.words {
		fmt.Println("checking word", word)
		// strs := strings.Split(s, " ")
		// for _, sub := range s {
		// for k := range set {
		// 	fmt.Println("checking k", k)
		// 	if !strings.Contains(k, word) {
		// 		fmt.Println("new word", word, k)
		// 		set[word] = struct{}{}
		// 		res = append(res, word)
		// 	}
		// }
		_, ok := set[word]
		if !ok {
			for k := range set {
				fmt.Println("checking k", k)
				if !strings.Contains(k, word) {
					fmt.Println("new word", word, k)
					set[word] = struct{}{}
					res = append(res, word)
				}
			}
		}

	}
	fmt.Printf("res %v\n", res)
	d.WordsDeduped = res
}
