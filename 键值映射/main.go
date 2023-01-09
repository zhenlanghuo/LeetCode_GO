package main

import "fmt"

func main() {
	ms := Constructor()
	ms.Insert("apple", 3)
	fmt.Println(ms.Sum("ap"))
	ms.Insert("app", 2)
	fmt.Println(ms.Sum("ap"))
}

type MapSum struct {
	root *TrieNode
	cnt  map[string]int
}

type TrieNode struct {
	val      int
	children [26]*TrieNode
}

func (n *TrieNode) Insert(word string, i int, val int) *TrieNode {
	node := n
	if node == nil {
		node = &TrieNode{
			//children: make([]*TrieNode, 26),
		}
	}

	node.val += val

	if i == len(word) {
		//node.val += val
		return node
	}

	node.children[word[i]-'a'] = node.children[word[i]-'a'].Insert(word, i+1, val)
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

//func (n *TrieNode) GetChildrenValSum() int {
//	if n == nil {
//		return 0
//	}
//
//	sum := n.val
//	for _, node := range n.children {
//		sum += node.GetChildrenValSum()
//	}
//	return sum
//}

func (n *TrieNode) GetValSumWithPrefix(prefix string) int {
	if n == nil {
		return 0
	}

	preFixNode := n.GetNode(prefix, 0)
	if preFixNode == nil {
		return 0
	}
	return preFixNode.val
}

func Constructor() MapSum {
	return MapSum{
		root: &TrieNode{},
		cnt:  make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	oldVal := this.cnt[key]
	this.root.Insert(key, 0, val-oldVal)
	this.cnt[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	return this.root.GetValSumWithPrefix(prefix)
}
