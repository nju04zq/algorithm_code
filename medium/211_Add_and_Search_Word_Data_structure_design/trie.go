package main

import "fmt"

type WordDictionary struct {
	root trieNode
}

type trieNode struct {
	exist    bool
	children [26]*trieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	p := &this.root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if p.children[j] == nil {
			p.children[j] = new(trieNode)
		}
		p = p.children[j]
	}
	p.exist = true
}

func (this *WordDictionary) SearchInternal(word string, p *trieNode) bool {
	if word == "" {
		return p.exist
	}
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			for _, child := range p.children {
				if child != nil && this.SearchInternal(word[i+1:], child) {
					return true
				}
			}
			return false
		}
		j := word[i] - 'a'
		if p.children[j] == nil {
			return false
		}
		p = p.children[j]
	}
	return p.exist
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.SearchInternal(word, &this.root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

func main() {
	t := Constructor()
	t.AddWord("bad")
	t.AddWord("dad")
	t.AddWord("mad")
	s := "pad"
	fmt.Printf("find %q, get %t\n", s, t.Search(s))
	s = "dad"
	fmt.Printf("find %q, get %t\n", s, t.Search(s))
	s = ".ad"
	fmt.Printf("find %q, get %t\n", s, t.Search(s))
	s = "b.."
	fmt.Printf("find %q, get %t\n", s, t.Search(s))
	s = ".adx"
	fmt.Printf("find %q, get %t\n", s, t.Search(s))
	s = ".ax"
	fmt.Printf("find %q, get %t\n", s, t.Search(s))
	s = "."
	fmt.Printf("find %q, get %t\n", s, t.Search(s))
	s = "d."
	fmt.Printf("find %q, get %t\n", s, t.Search(s))
}
