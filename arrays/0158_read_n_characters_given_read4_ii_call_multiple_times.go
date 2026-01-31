package arrays


/*
Difficulty: Hard
Tags: [Add relevant tags]
Companies: [Add company names]
*/

import "leetcode/utils"

// 158. Read N Characters Given Read4 II - Call multiple times
// https://leetcode.com/problems/read-n-characters-given-read4-ii-call-multiple-times/

// read4_158 is a mock API that reads 4 characters from a file into buf4
// This is provided by the system, we just need to implement the read function
// Renamed to avoid conflict with 0157
func read4_158(buf4 []byte) int {
	// This is a mock implementation - the actual implementation is provided by the system
	// In real LeetCode environment, this function is provided
	return 0
}

// Solution is the main struct that maintains state between multiple read calls
type Solution struct {
	buffer      []byte // Internal buffer to store leftover characters
	bufferIndex int    // Current index in the buffer
	bufferSize  int    // Total valid characters in the buffer
}

// Constructor_158 creates a new Solution instance
// Renamed to avoid conflict with 0146
func Constructor_158() Solution {
	return Solution{
		buffer:      make([]byte, 4),
		bufferIndex: 0,
		bufferSize:  0,
	}
}

// read reads n characters from a file into buf using the read4 API
// This function can be called multiple times and maintains state between calls
// Time Complexity: O(n), Space Complexity: O(1)
func (s *Solution) read(buf []byte, n int) int {
	totalRead := 0

	// First, read any leftover characters from the buffer
	for totalRead < n && s.bufferIndex < s.bufferSize {
		buf[totalRead] = s.buffer[s.bufferIndex]
		totalRead++
		s.bufferIndex++
	}

	// If we still need more characters, read from the file
	for totalRead < n {
		// Read up to 4 characters using read4
		readCount := read4_158(s.buffer)

		if readCount == 0 {
			// No more characters to read
			break
		}

		// Reset buffer index and size
		s.bufferIndex = 0
		s.bufferSize = readCount

		// Copy characters from buffer to buf
		bytesToCopy := utils.Min(readCount, n-totalRead)
		for i := 0; i < bytesToCopy; i++ {
			buf[totalRead] = s.buffer[i]
			totalRead++
			s.bufferIndex++
		}

		// If we couldn't copy all bytes from read4, we have leftovers
		if bytesToCopy < readCount {
			// Update buffer index to point to the first leftover character
			s.bufferIndex = bytesToCopy
			break
		}
	}

	return totalRead
}

// readOptimized is an optimized version with better memory management
// Time Complexity: O(n), Space Complexity: O(1)
func (s *Solution) readOptimized(buf []byte, n int) int {
	totalRead := 0

	// Read from buffer first if we have leftovers
	for totalRead < n && s.bufferIndex < s.bufferSize {
		buf[totalRead] = s.buffer[s.bufferIndex]
		totalRead++
		s.bufferIndex++
	}

	// Read directly from file if we need more
	for totalRead < n {
		// Read directly into the destination buffer if possible
		remaining := n - totalRead
		if remaining >= 4 {
			// We have space for at least 4 characters, read directly
			readCount := read4_158(buf[totalRead:])
			totalRead += readCount
			if readCount < 4 {
				// No more characters to read
				break
			}
		} else {
			// We need to use the buffer because we don't have space for 4 characters
			readCount := read4_158(s.buffer)
			if readCount == 0 {
				break
			}

			// Copy what we can
			bytesToCopy := utils.Min(readCount, remaining)
			for i := 0; i < bytesToCopy; i++ {
				buf[totalRead] = s.buffer[i]
				totalRead++
			}

			// Store leftovers in buffer
			if bytesToCopy < readCount {
				s.bufferIndex = bytesToCopy
				s.bufferSize = readCount
			} else {
				// Reset buffer since we used all characters
				s.bufferIndex = 0
				s.bufferSize = 0
			}
		}
	}

	return totalRead
}

// readWithCircularBuffer uses a circular buffer approach for better efficiency
// Time Complexity: O(n), Space Complexity: O(1)
func (s *Solution) readWithCircularBuffer(buf []byte, n int) int {
	totalRead := 0

	// Read from circular buffer first
	for totalRead < n && s.bufferIndex < s.bufferSize {
		buf[totalRead] = s.buffer[s.bufferIndex]
		totalRead++
		s.bufferIndex = (s.bufferIndex + 1) % 4
		if s.bufferIndex == 0 {
			s.bufferSize = 0
		}
	}

	// Read from file if needed
	for totalRead < n {
		// Read into buffer
		readCount := read4_158(s.buffer)
		if readCount == 0 {
			break
		}

		// Copy to destination
		bytesToCopy := utils.Min(readCount, n-totalRead)
		for i := 0; i < bytesToCopy; i++ {
			buf[totalRead] = s.buffer[i]
			totalRead++
		}

		// Store leftovers
		if bytesToCopy < readCount {
			s.bufferIndex = bytesToCopy
			s.bufferSize = readCount
		} else {
			s.bufferIndex = 0
			s.bufferSize = 0
		}
	}

	return totalRead
}

// readUsingTwoPointers uses two-pointer technique for efficient copying
// Time Complexity: O(n), Space Complexity: O(1)
func (s *Solution) readUsingTwoPointers(buf []byte, n int) int {
	totalRead := 0
	bufPtr := 0

	// Read from internal buffer first
	for totalRead < n && s.bufferIndex < s.bufferSize {
		buf[bufPtr] = s.buffer[s.bufferIndex]
		bufPtr++
		totalRead++
		s.bufferIndex++
	}

	// Read from file
	for totalRead < n {
		// Read into internal buffer
		readCount := read4_158(s.buffer)
		if readCount == 0 {
			break
		}

		// Reset buffer index
		s.bufferIndex = 0
		s.bufferSize = readCount

		// Copy as much as we can
		bytesToCopy := utils.Min(readCount, n-totalRead)
		for i := 0; i < bytesToCopy; i++ {
			buf[bufPtr] = s.buffer[i]
			bufPtr++
			totalRead++
			s.bufferIndex++
		}

		// If we have leftovers, break
		if bytesToCopy < readCount {
			break
		}
	}

	return totalRead
}

// readBatch processes characters in batches for better performance
// Time Complexity: O(n), Space Complexity: O(1)
func (s *Solution) readBatch(buf []byte, n int) int {
	totalRead := 0

	// Process buffer leftovers
	if s.bufferIndex < s.bufferSize {
		available := s.bufferSize - s.bufferIndex
		toCopy := utils.Min(available, n)
		copy(buf[:toCopy], s.buffer[s.bufferIndex:s.bufferIndex+toCopy])
		totalRead += toCopy
		s.bufferIndex += toCopy
	}

	// Process in batches of 4
	for totalRead < n {
		remaining := n - totalRead

		if remaining >= 4 {
			// Read directly into destination
			readCount := read4_158(buf[totalRead:])
			totalRead += readCount
			if readCount < 4 {
				break
			}
		} else {
			// Use buffer for small remaining amounts
			readCount := read4_158(s.buffer)
			if readCount == 0 {
				break
			}

			toCopy := utils.Min(readCount, remaining)
			copy(buf[totalRead:totalRead+toCopy], s.buffer[:toCopy])
			totalRead += toCopy

			// Store leftovers
			if toCopy < readCount {
				s.bufferIndex = toCopy
				s.bufferSize = readCount
			} else {
				s.bufferIndex = 0
				s.bufferSize = 0
			}
		}
	}

	return totalRead
}