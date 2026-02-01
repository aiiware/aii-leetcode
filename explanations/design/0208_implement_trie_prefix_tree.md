# 0208 - Implement Trie (Prefix Tree)

## Problem Statement
Implement a trie (prefix tree) with the following operations:
- `insert(word)`: Inserts the string `word` into the trie
- `search(word)`: Returns `true` if the string `word` is in the trie, `false` otherwise
- `startsWith(prefix)`: Returns `true` if there is any word in the trie that starts with the given `prefix`, `false` otherwise

**Example:**
```
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
```

## Solution Approach
A **trie** (pronounced "try") is a tree-like data structure that stores strings character by character. Each node represents a character, and paths from the root to nodes with end markers represent complete words.

### Key Insight
Tries are particularly efficient for prefix-based operations because:
1. **Prefix search**: O(L) where L is prefix length
2. **Word search**: O(L) where L is word length  
3. **Memory efficient**: Shared prefixes are stored only once

## Algorithm Design

### Data Structures
- **TrieNode**: Each node contains:
  - `children`: Array/map of child nodes (26 for lowercase English letters)
  - `isEnd`: Boolean flag indicating if this node completes a word

### `insert(word)`
1. Start at the root
2. For each character in the word:
   - Get the index: `char - 'a'`
   - If child doesn't exist, create a new TrieNode
   - Move to the child node
3. Mark the final node as end of word

### `search(word)`
1. Start at the root
2. For each character in the word:
   - Get the index: `char - 'a'`
   - If child doesn't exist, return `false`
   - Move to the child node
3. Return `true` if final node is marked as end of word

### `startsWith(prefix)`
1. Start at the root
2. For each character in the prefix:
   - Get the index: `char - 'a'`
   - If child doesn't exist, return `false`
   - Move to the child node
3. Return `true` (we successfully traversed the entire prefix)

## Complexity Analysis

### Time Complexity
- **`insert(word)`**: O(L) where L is the length of the word
- **`search(word)`**: O(L) where L is the length of the word
- **`startsWith(prefix)`**: O(P) where P is the length of the prefix

### Space Complexity
- **Worst case**: O(N × L) where N is number of words and L is average word length
- **Best case (shared prefixes)**: Much less than worst case due to prefix sharing

## Implementation Details

### Go Implementation
```go
type Trie struct {
    children [26]*Trie
    isEnd    bool
}

func Constructor() Trie {
    return Trie{}
}

func (this *Trie) Insert(word string) {
    node := this
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            node.children[idx] = &Trie{}
        }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (this *Trie) Search(word string) bool {
    node := this
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return false
        }
        node = node.children[idx]
    }
    return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
    node := this
    for _, ch := range prefix {
        idx := ch - 'a'
        if node.children[idx] == nil {
            return false
        }
        node = node.children[idx]
    }
    return true
}
```

### Key Points
1. **Array-based children**: Fixed size array (26) for lowercase English letters
2. **Memory optimization**: Could use map for Unicode support or sparse data
3. **End marker**: `isEnd` flag distinguishes between prefix and complete word
4. **Root node**: Empty node that doesn't represent any character

## Test Cases

### Example 1
```
Operations:
- insert("apple")
- search("apple")   → true
- search("app")     → false
- startsWith("app") → true
- insert("app")
- search("app")     → true
```

### Example 2
```
Operations:
- insert("cat")
- insert("car")
- insert("card")
- search("cat")     → true
- search("car")     → true
- search("card")    → true
- search("ca")      → false
- startsWith("ca")  → true
- startsWith("car") → true
- startsWith("cad") → false
```

### Edge Cases
1. **Empty string**: Can be inserted and searched (root's `isEnd` would be true)
2. **Duplicate insertion**: Inserting same word twice doesn't break the structure
3. **Non-existent search**: Returns false without errors
4. **Case sensitivity**: Assumes lowercase English letters only

## Related Problems
- **0211 - Design Add and Search Words Data Structure**: Trie with wildcard support
- **0677 - Map Sum Pairs**: Trie with value accumulation
- **0648 - Replace Words**: Trie for prefix replacement
- **0212 - Word Search II**: Trie for word search in grid

## Learning Points
1. **Prefix trees**: Understanding the trie data structure
2. **Character indexing**: Converting characters to array indices
3. **Space-time tradeoff**: Fixed array vs. map for children storage
4. **Prefix operations**: Efficient prefix matching algorithms

## Real-World Applications
- **Autocomplete systems**: Search engines, IDEs, mobile keyboards
- **Spell checkers**: Dictionary lookups and suggestions
- **IP routing tables**: Longest prefix matching in networking
- **Genome sequencing**: Storing and searching DNA sequences
- **Contact search**: Phone book applications with prefix search