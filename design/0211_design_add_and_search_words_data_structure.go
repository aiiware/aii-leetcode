package design

/*
211. Design Add and Search Words Data Structure

Design a data structure that supports adding new words and finding if a string matches any 
previously added string.

Implement the WordDictionary class:
- WordDictionary() Initializes the object.
- void addWord(word) Adds word to the data structure, it can be matched later.
- bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.

Example:
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True

Constraints:
- 1 <= word.length <= 25
- word in addWord consists of lowercase English letters.
- word in search consist of '.' or lowercase English letters.
- There will be at most 10^4 calls in total to addWord and search.
*/

/*
Difficulty: Medium
Tags: String, Depth-First Search, Design, Trie
Companies: Amazon, Facebook, Google, Microsoft, Apple, Bloomberg, Uber
*/

// TrieNode for WordDictionary
type WordTrieNode struct {
    children map[byte]*WordTrieNode
    isEnd    bool
}

func NewWordTrieNode() *WordTrieNode {
    return &WordTrieNode{
        children: make(map[byte]*WordTrieNode),
        isEnd:    false,
    }
}

// WordDictionary implements a dictionary with wildcard search
type WordDictionary struct {
    root *WordTrieNode
}

func ConstructorWordDictionary() WordDictionary {
    return WordDictionary{
        root: NewWordTrieNode(),
    }
}

func (this *WordDictionary) AddWord(word string) {
    node := this.root
    for i := 0; i < len(word); i++ {
        ch := word[i]
        if node.children[ch] == nil {
            node.children[ch] = NewWordTrieNode()
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.searchDFS(this.root, word, 0)
}

// searchDFS performs DFS search with wildcard support
func (this *WordDictionary) searchDFS(node *WordTrieNode, word string, index int) bool {
    if node == nil {
        return false
    }
    
    if index == len(word) {
        return node.isEnd
    }
    
    ch := word[index]
    if ch == '.' {
        // Try all possible children
        for _, child := range node.children {
            if this.searchDFS(child, word, index+1) {
                return true
            }
        }
        return false
    } else {
        // Regular character match
        child := node.children[ch]
        if child == nil {
            return false
        }
        return this.searchDFS(child, word, index+1)
    }
}

// Array-based implementation (faster for lowercase letters)
type WordDictionaryArray struct {
    children [26]*WordDictionaryArray
    isEnd    bool
}

func ConstructorWordDictionaryArray() WordDictionaryArray {
    return WordDictionaryArray{}
}

func (this *WordDictionaryArray) AddWord(word string) {
    node := this
    for i := 0; i < len(word); i++ {
        idx := word[i] - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &WordDictionaryArray{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (this *WordDictionaryArray) Search(word string) bool {
    return this.searchDFSArray(this, word, 0)
}

func (this *WordDictionaryArray) searchDFSArray(node *WordDictionaryArray, word string, index int) bool {
    if node == nil {
        return false
    }
    
    if index == len(word) {
        return node.isEnd
    }
    
    ch := word[index]
    if ch == '.' {
        // Try all 26 possible children
        for i := 0; i < 26; i++ {
            if node.children[i] != nil && this.searchDFSArray(node.children[i], word, index+1) {
                return true
            }
        }
        return false
    } else {
        idx := ch - 'a'
        child := node.children[idx]
        if child == nil {
            return false
        }
        return this.searchDFSArray(child, word, index+1)
    }
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */