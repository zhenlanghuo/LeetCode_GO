package main

import "fmt"

func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
	fmt.Println(t.Search("app"))
	fmt.Println(t.StartsWith("app"))
	t.Insert("app")
	fmt.Println(t.Search("app"))
}

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	val      bool
	children [26]*TrieNode
}

func Constructor() Trie {
	return Trie{
		root: nil,
	}
}

func (n *TrieNode) Insert(word string, i int) *TrieNode {
	node := n
	if node == nil {
		node = &TrieNode{
			//children: make([]*TrieNode, 26),
		}
	}

	if i == len(word) {
		node.val = true
		return node
	}

	node.children[word[i]-'a'] = node.children[word[i]-'a'].Insert(word, i+1)
	return node
}

func (n *TrieNode) GetNode(word string, i int) *TrieNode {
	if i == len(word) {
		return n
	}

	if n == nil {
		return nil
	}

	return n.children[word[i]-'a'].GetNode(word, i+1)
}

func (this *Trie) Insert(word string) {
	this.root = this.root.Insert(word, 0)
}

func (this *Trie) Search(word string) bool {
	node := this.root.GetNode(word, 0)
	if node != nil && node.val {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root.GetNode(prefix, 0)
	if node != nil {
		return true
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
