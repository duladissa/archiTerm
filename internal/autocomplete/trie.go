package autocomplete

// TrieNode represents a node in the trie
type TrieNode struct {
	children    map[rune]*TrieNode
	isEnd       bool
	command     string
	description string
}

// Trie is a prefix tree for fast command lookup
type Trie struct {
	root *TrieNode
}

// NewTrie creates a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

// Insert adds a command to the trie
func (t *Trie) Insert(command, description string) {
	node := t.root
	for _, ch := range command {
		if node.children[ch] == nil {
			node.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.command = command
	node.description = description
}

// Search finds all commands with the given prefix
func (t *Trie) Search(prefix string) []Match {
	node := t.root
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return t.collectAll(node, prefix)
}

// collectAll collects all commands under a node
func (t *Trie) collectAll(node *TrieNode, prefix string) []Match {
	var results []Match

	if node.isEnd {
		results = append(results, Match{
			Command:     node.command,
			Description: node.description,
			Score:       100, // Exact prefix match gets high score
		})
	}

	for ch, child := range node.children {
		childResults := t.collectAll(child, prefix+string(ch))
		results = append(results, childResults...)
	}

	return results
}

// GetCompletion returns the completion suffix for a prefix
func (t *Trie) GetCompletion(prefix string) string {
	node := t.root
	for _, ch := range prefix {
		if node.children[ch] == nil {
			return ""
		}
		node = node.children[ch]
	}

	// Find the first complete command from this point
	return t.findFirstCompletion(node, "")
}

// findFirstCompletion finds the shortest completion
func (t *Trie) findFirstCompletion(node *TrieNode, suffix string) string {
	if node.isEnd {
		return suffix
	}

	// Take the first available path (alphabetically)
	var minRune rune
	var minChild *TrieNode
	for ch, child := range node.children {
		if minChild == nil || ch < minRune {
			minRune = ch
			minChild = child
		}
	}

	if minChild == nil {
		return suffix
	}

	return t.findFirstCompletion(minChild, suffix+string(minRune))
}
