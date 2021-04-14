package main

import "fmt"

type Trie struct {
	Root *Node
}

type Node struct {
	Next map[rune] *Node
	IsWord bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Root: &Node{
			Next: make(map[rune] *Node),
		},
	}
}


/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	cur := t.Root
	for _, c := range word {
		fmt.Print(c)
		_, ok := cur.Next[c]
		if !ok {
			cur.Next[c] = &Node{
				Next: make(map[rune] *Node),
			}
		}

		cur = cur.Next[c]
	}

	if !cur.IsWord {
		cur.IsWord = true
	}

}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	cur := t.Root
	for _, c := range word {
		_, ok := cur.Next[c]
		if !ok {
			return false
		}

		cur = cur.Next[c]
	}

	return cur.IsWord
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	cur := t.Root
	for _, c := range prefix {
		_, ok := cur.Next[c]
		if !ok {
			return false
		}

		cur = cur.Next[c]
	}

	return true
}

func main() {
	words := [] string {"test", "tree", "tri", "triple"}
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie)

	searchTestWords := [] string {"test", "time", "throw"}
	var searchTestRes [] bool
	for _, word := range searchTestWords {
		searchTestRes = append(searchTestRes, trie.Search(word))
	}

	fmt.Println(searchTestRes)
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */