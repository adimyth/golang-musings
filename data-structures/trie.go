package datastructures

type TrieNode struct {
	children [26]*TrieNode
	wordEnd  bool
}

type Trie struct {
	rootNode *TrieNode
}

// Constructor to initialize the Trie
func Constructor() Trie {
	return Trie{rootNode: &TrieNode{}}
}

// Insert a word into the Trie
func (t *Trie) Insert(word string) {
	// Start from the root node
	node := t.rootNode

	// Insert each character of the word
	for _, char := range word {
		index := char - 'a'
		// Create a new node if the character is not found
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		// Move to the next node
		node = node.children[index]
	}
	node.wordEnd = true
}

// Search for a word in the Trie
func (t *Trie) Search(word string) bool {
	// Start from the root node
	node := t.rootNode

	// Traverse the Trie following the characters of the word one by one
	for _, char := range word {
		index := char - 'a'
		// Return false if the character is not found
		if node.children[index] == nil {
			return false
		}
		// Otherwise move to the next node in the Trie
		node = node.children[index]
	}
	// If at the end of the word, return true if the node is marked as end of word
	return node.wordEnd
}
