package chapter17

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()
	trie.Insert("ace")
	trie.Insert("bad")
	trie.Insert("cat")
	trie.Insert("bat")
	trie.Insert("batter")
	result := trie.Search("cat")
	assert.Equal(t, &TrieNode{Children: map[rune]*TrieNode{
		'*': nil,
	}}, result)
	result = trie.Search("nope")
	assert.Nil(t, result)
}

func TestTrieCollectAllWords(t *testing.T) {
	trie := NewTrie()
	trie.Insert("ace")
	trie.Insert("bad")
	trie.Insert("cat")
	trie.Insert("bat")
	trie.Insert("batter")
	result := trie.CollectAllWords("", nil, []string{})
	sort.Strings(result) // we need to sort because we are dealing with a map
	assert.Equal(t, []string{"ace", "bad", "bat", "batter", "cat"}, result)
}

func TestTrieAutoComplete(t *testing.T) {
	trie := NewTrie()
	trie.Insert("ace")
	trie.Insert("bad")
	trie.Insert("cat")
	trie.Insert("bat")
	trie.Insert("batter")
	result := trie.AutoComplete("ba")
	sort.Strings(result) // we need to sort because we are dealing with a map
	// These are the characters or words that an autocomplete feature would autocomplete or
	// offer up as a word. The calling code must insert the prefix if it wants to display
	// full words.
	assert.Equal(t, []string{"d", "t", "tter"}, result)
}

func TestTrieAutoCorrect(t *testing.T) {
	trie := NewTrie()
	trie.Insert("cat")
	trie.Insert("catnap")
	trie.Insert("catnip")
	result := trie.AutoCorrect("catnar")
	assert.Equal(t, "catnap", result)
	result = trie.AutoCorrect("calkka")
	assert.Equal(t, "cat", result)
}

func TestTraverse(t *testing.T) {
	trie := NewTrie()
	trie.Insert("ace")
	trie.Insert("bad")
	trie.Insert("cat")
	ch := Traverse(trie.Root, []rune{})
	assert.Equal(t, "ace*bad*cat*", string(ch))
}
