package dedupe

type Data struct {
	Name  string
	Words []string

	Map map[string]int
}

func NewSet() *Data {
	return &Data{
		Map: make(map[string]int),
	}
}

func (d *Data) Add(word string) {
	d.Words = append(d.Words, word)
	d.Map[word] += 1
}
