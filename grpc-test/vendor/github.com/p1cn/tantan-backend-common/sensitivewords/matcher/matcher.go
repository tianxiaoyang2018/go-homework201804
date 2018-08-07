package matcher

type Matcher interface {
	Match(in string) bool
	Replace(str string, cover string) (string, int)
}

type Word struct {
	Content string
	Flags   uint8
}

type Dictionary struct {
	Words []*Word
}

func (d *Dictionary) AddWord(content string, flags uint8) {
	d.Words = append(d.Words, &Word{Content: content, Flags: flags})
}
