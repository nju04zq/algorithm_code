package main

import "fmt"

type TrieNode struct {
	Exist    bool
	Children [26]*TrieNode
}

type Trie struct {
	Root TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := &this.Root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		if node.Children[j] == nil {
			node.Children[j] = new(TrieNode)
		}
		node = node.Children[j]
	}
	node.Exist = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := &this.Root
	for i := 0; i < len(word); i++ {
		j := word[i] - 'a'
		node = node.Children[j]
		if node == nil {
			return false
		}
	}
	return node.Exist
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := &this.Root
	for i := 0; i < len(prefix); i++ {
		j := prefix[i] - 'a'
		node = node.Children[j]
		if node == nil {
			return false
		}
	}
	if node.Exist {
		return true
	}
	for _, child := range node.Children {
		if child != nil {
			return true
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	t := Constructor()
	t.Insert("aba")
	t.Insert("abc")
	t.Insert("abcd")
	s := "abcd"
	fmt.Printf("%q exist %t\n", s, t.Search(s))
	s = "abd"
	fmt.Printf("%q exist %t\n", s, t.Search(s))
	s = "ab"
	fmt.Printf("prefix %q exist %t\n", s, t.StartsWith(s))
	s = "abd"
	fmt.Printf("prefix %q exist %t\n", s, t.StartsWith(s))
}
