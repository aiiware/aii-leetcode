package design

/*
208. Implement Trie (Prefix Tree)

A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently 
store and retrieve keys in a dataset of strings. There are various applications of this 
data structure, such as autocomplete and spellchecker.

Implement the Trie class:
- Trie() Initializes the trie object.
- void insert(String word) Inserts the string word into the trie.
- boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
- boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.

Example 1:
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

Constraints:
- 1 <= word.length, prefix.length <= 2000
- word and prefix consist only of lowercase English letters.
- At most 3 * 10^4 calls in total will be made to insert, search, and startsWith.
*/

/*
Difficulty: Medium
Tags: Hash Table, String, Design, Trie
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber, Oracle, TikTok, LinkedIn
*/

// TrieNode represents a node in the Trie
type TrieNode struct {
    children map[byte]*TrieNode
    isEnd    bool
}

// NewTrieNode creates a new TrieNode
func NewTrieNode() *TrieNode {
    return &TrieNode{
        children: make(map[byte]*TrieNode),
        isEnd:    false,
    }
}

// Trie implements a prefix tree
type Trie struct {
    root *TrieNode
}

// ConstructorTrie initializes your data structure here
func ConstructorTrie() Trie {
    return Trie{
        root: NewTrieNode(),
    }
}

// Insert inserts a word into the trie
func (this *Trie) Insert(word string) {
    node := this.root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if node.children[ch] == nil {
            node.children[ch] = NewTrieNode()
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

// Search returns if the word is in the trie
func (this *Trie) Search(word string) bool {
    node := this.searchPrefix(word)
    return node != nil && node.isEnd
}

// StartsWith returns if there is any word in the trie that starts with the given prefix
func (this *Trie) StartsWith(prefix string) bool {
    node := this.searchPrefix(prefix)
    return node != nil
}

// searchPrefix returns the node where the prefix ends, or nil if prefix doesn't exist
func (this *Trie) searchPrefix(prefix string) *TrieNode {
    node := this.root
    for i := 0; i < len(prefix); i++ {
        ch := prefix[i]
        if node.children[ch] == nil {
            return nil
        }
        node = node.children[ch]
    }
    return node
}

// Alternative implementation using array instead of map (faster for lowercase letters only)
type TrieArray struct {
    children [26]*TrieArray
    isEnd    bool
}

func ConstructorArray() TrieArray {
    return TrieArray{}
}

func (this *TrieArray) Insert(word string) {
    node := this
    for i := 0; i < len(word); i++ {
        idx := word[i] - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &TrieArray{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (this *TrieArray) Search(word string) bool {
    node := this.searchPrefix(word)
    return node != nil && node.isEnd
}

func (this *TrieArray) StartsWith(prefix string) bool {
    node := this.searchPrefix(prefix)
    return node != nil
}

func (this *TrieArray) searchPrefix(prefix string) *TrieArray {
    node := this
    for i := 0; i < len(prefix); i++ {
        idx := prefix[i] - 'a'
        if node.children[idx] == nil {
            return nil
        }
        node = node.children[idx]
    }
    return node
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */