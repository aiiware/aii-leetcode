package leetcode

// FullJustify solves LeetCode problem 0068: Text Justification
// Difficulty: Hard
// Tags: String, Simulation
//
// Given an array of strings words and a width maxWidth, format the text such that each line has
// exactly maxWidth characters and is fully (left and right) justified.
//
// You should pack your words in a greedy approach; that is, pack as many words as you can in each line.
// Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.
//
// Extra spaces between words should be distributed as evenly as possible. If the number of spaces
// on a line does not divide evenly between words, the empty slots on the left will be assigned
// more spaces than the slots on the right.
//
// For the last line of text, it should be left-justified, and no extra space is inserted between words.
//
// Example 1:
// Input: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
// Output:
// [
//    "This    is    an",
//    "example  of text",
//    "justification.  "
// ]
//
// Example 2:
// Input: words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
// Output:
// [
//   "What   must   be",
//   "acknowledgment  ",
//   "shall be        "
// ]
//
// Time complexity: O(n), Space complexity: O(n)
func FullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return []string{}
	}

	result := []string{}
	lineStart := 0
	currentLength := 0

	for i := 0; i < len(words); i++ {
		word := words[i]
		
		// Calculate if we can add this word to current line
		// currentLength is sum of lengths of words already in line
		// (i - lineStart) is number of spaces needed between words
		if currentLength+len(word)+(i-lineStart) > maxWidth {
			// Can't add this word, justify current line
			result = append(result, justifyLine(words[lineStart:i], currentLength, maxWidth, false))
			lineStart = i
			currentLength = 0
		}
		
		currentLength += len(word)
	}

	// Handle last line (left-justified)
	result = append(result, justifyLine(words[lineStart:], currentLength, maxWidth, true))
	return result
}

// justifyLine justifies a line of words
// words: words in the line
// wordsLength: total length of all words (without spaces)
// maxWidth: target width of the line
// isLastLine: whether this is the last line (left-justified)
func justifyLine(words []string, wordsLength, maxWidth int, isLastLine bool) string {
	if len(words) == 0 {
		return ""
	}

	if len(words) == 1 || isLastLine {
		// Single word or last line: left-justified
		line := words[0]
		for i := 1; i < len(words); i++ {
			line += " " + words[i]
		}
		// Add trailing spaces
		for len(line) < maxWidth {
			line += " "
		}
		return line
	}

	// Multiple words and not last line: fully justified
	totalSpaces := maxWidth - wordsLength
	gaps := len(words) - 1
	
	// Calculate base spaces per gap and extra spaces
	baseSpaces := totalSpaces / gaps
	extraSpaces := totalSpaces % gaps

	line := words[0]
	for i := 1; i < len(words); i++ {
		// Add base spaces
		for s := 0; s < baseSpaces; s++ {
			line += " "
		}
		// Add extra space for first 'extraSpaces' gaps
		if i <= extraSpaces {
			line += " "
		}
		line += words[i]
	}

	return line
}