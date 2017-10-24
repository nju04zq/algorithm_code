package main

import "fmt"
import "bytes"

type line struct {
	buf      *bytes.Buffer
	maxWidth int
	words    []string
	total    int
	wordsLen int
}

func (l *line) init(maxWidth int) *line {
	l.maxWidth = maxWidth
	l.buf = bytes.NewBuffer(nil)
	l.words = make([]string, 0)
	return l
}

func (l *line) add(word string) bool {
	total := len(word) + l.total
	if total > l.maxWidth {
		return false
	}
	l.words = append(l.words, word)
	l.total = (total + 1)
	l.wordsLen += len(word)
	return true
}

func (l *line) reset() {
	l.buf.Reset()
	l.words = l.words[:0]
	l.total = 0
	l.wordsLen = 0
}

func (l *line) format() string {
	spaces := l.maxWidth - l.wordsLen
	var base, mod int
	if len(l.words) == 1 {
		l.buf.WriteString(l.words[0])
		for i := 0; i < spaces; i++ {
			l.buf.WriteByte(' ')
		}
		s := l.buf.String()
		l.reset()
		return s
	}
	base = spaces / (len(l.words) - 1)
	mod = spaces % (len(l.words) - 1)
	for i, word := range l.words {
		l.buf.WriteString(word)
		if i == len(l.words)-1 {
			break
		}
		if mod > 0 {
			spaces = base + 1
		} else {
			spaces = base
		}
		for i := 0; i < spaces; i++ {
			l.buf.WriteByte(' ')
		}
		mod--
	}
	s := l.buf.String()
	l.reset()
	return s
}

func (l *line) formatLastLine() string {
	spaces := l.maxWidth - l.wordsLen
	for _, word := range l.words {
		l.buf.WriteString(word)
		if spaces > 0 {
			l.buf.WriteByte(' ')
			spaces--
		}
	}
	for i := 0; i < spaces; i++ {
		l.buf.WriteByte(' ')
	}
	s := l.buf.String()
	return s
}

func fullJustify(words []string, maxWidth int) []string {
	res := make([]string, 0)
	l := new(line).init(maxWidth)
	for _, word := range words {
		if !l.add(word) {
			res = append(res, l.format())
			l.add(word)
		}
	}
	res = append(res, l.formatLastLine())
	return res
}

func testJustify(words []string, maxWidth int) {
	fmt.Printf("words %v, maxWidth %d\n", words, maxWidth)
	res := fullJustify(words, maxWidth)
	for _, s := range res {
		fmt.Printf("%q\n", s)
	}
}

func main() {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	width := 16
	testJustify(words, width)
	words = []string{"This"}
	width = 10
	testJustify(words, width)
}
