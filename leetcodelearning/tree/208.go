package tree

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curNode := this
	for i := 0; i < len(word); i++ {
		letterIndex := word[i] - 'a'
		if curNode.children[letterIndex] == nil {
			curNode.children[letterIndex] = &Trie{}
		}
		if i == len(word)-1 {
			curNode.children[letterIndex].isEnd = true
		}
		curNode = curNode.children[letterIndex]
	}
}

func (this *Trie) findLastLetterNode(word string) *Trie {
	curNode := this
	for i := 0; i < len(word); i++ {
		letterIndex := word[i] - 'a'
		if curNode.children[letterIndex] == nil {
			return nil
		}
		curNode = curNode.children[letterIndex]
	}
	return curNode
}

func (this *Trie) Search(word string) bool {
	lastNode := this.findLastLetterNode(word)
	return lastNode != nil && lastNode.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	lastNode := this.findLastLetterNode(prefix)
	return lastNode != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
