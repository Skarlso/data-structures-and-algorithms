package chapter17

type TrieNode struct {
	Children map[rune]*TrieNode
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Search(word string) *TrieNode {
	// Set root to current node.
	currentNode := t.Root

	for _, c := range word {
		if v, ok := currentNode.Children[c]; ok {
			currentNode = v
			continue
		}

		// We didn't find our word.
		return nil
	}

	// We found our word and return the current node to help the insert later on.
	return currentNode
}

func (t *Trie) Insert(word string) *TrieNode {
	// Set root to current node.
	currentNode := t.Root

	for _, c := range word {
		// Update the node and go on...
		if v, ok := currentNode.Children[c]; ok {
			currentNode = v
			continue
		}

		// Create a new node with this new character
		next := &TrieNode{Children: make(map[rune]*TrieNode)}
		currentNode.Children[c] = next
		currentNode = next
	}

	// Lastly, when we are done finally, insert a last character to the last node.
	currentNode.Children['*'] = nil
	return currentNode
}

// CollectAllWords will collect all available words from a given Child Node.
// This is where it becomes handy that Search returns the last node.
// Because we'll search from that node onward for all available words in our
// autocomplete feature.
func (t *Trie) CollectAllWords(word string, node *TrieNode, words []string) []string {
	currentNode := t.Root
	if node != nil {
		currentNode = node
	}

	for k, v := range currentNode.Children {
		if k == '*' {
			words = append(words, word)
			continue
		}

		words = t.CollectAllWords(word+string(k), v, words)
	}

	return words
}

func (t *Trie) AutoComplete(prefix string) []string {
	currentNode := t.Search(prefix)
	if currentNode == nil {
		return nil
	}
	return t.CollectAllWords("", currentNode, []string{})
}

func (t *Trie) AutoCorrect(prefix string) string {
	node := t.Root
	wordFoundSoFar := ""
	for _, c := range prefix {
		if v, ok := node.Children[c]; ok {
			wordFoundSoFar += string(c)
			node = v
			continue
		}

		return wordFoundSoFar + t.CollectAllWords("", node, []string{})[0]
	}

	return prefix
}

func Traverse(node *TrieNode, characters []rune) []rune {
	if node == nil {
		return characters
	}

	for k, v := range node.Children {
		characters = append(characters, k)
		characters = Traverse(v, characters)
	}
	return characters
}
