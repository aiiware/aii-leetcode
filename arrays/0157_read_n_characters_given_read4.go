package arrays

import "leetcode/utils"

// 157. Read N Characters Given Read4
// https://leetcode.com/problems/read-n-characters-given-read4/

// read4 is a mock API that reads 4 characters from a file into buf4
// This is provided by the system, we just need to implement the read function
func read4(buf4 []byte) int {
	// This is a mock implementation - the actual implementation is provided by the system
	// In real LeetCode environment, this function is provided
	return 0
}

// read reads n characters from a file into buf using the read4 API
// Time Complexity: O(n), Space Complexity: O(1)
func read(buf []byte, n int) int {
	totalRead := 0
	buf4 := make([]byte, 4)

	for totalRead < n {
		// Read up to 4 characters using read4
		readCount := read4(buf4)

		if readCount == 0 {
			// No more characters to read
			break
		}

		// Copy characters from buf4 to buf
		bytesToCopy := utils.Min(readCount, n-totalRead)
		copy(buf[totalRead:], buf4[:bytesToCopy])
		totalRead += bytesToCopy

		if bytesToCopy < readCount {
			// We've read all n characters, but read4 returned more
			break
		}
	}

	return totalRead
}

// readOptimized is an optimized version with better memory usage
// Time Complexity: O(n), Space Complexity: O(1)
func readOptimized(buf []byte, n int) int {
	totalRead := 0
	buf4 := make([]byte, 4)

	for totalRead < n {
		// Read using read4
		readCount := read4(buf4)

		// If nothing was read, we're done
		if readCount == 0 {
			break
		}

		// Calculate how many bytes we can copy
		bytesToCopy := utils.Min(readCount, n-totalRead)

		// Copy the bytes
		for i := 0; i < bytesToCopy; i++ {
			buf[totalRead] = buf4[i]
			totalRead++
		}

		// If we couldn't copy all bytes from read4, we're done
		if bytesToCopy < readCount {
			break
		}
	}

	return totalRead
}

// readWithBuffer is another approach that tracks buffer state between calls
// This is useful when read might be called multiple times
type Reader struct {
	buf4      []byte
	buf4Index int
	buf4Size  int
}

func NewReader() *Reader {
	return &Reader{
		buf4:      make([]byte, 4),
		buf4Index: 0,
		buf4Size:  0,
	}
}

// readWithBuffer reads n characters using buffered approach
// Time Complexity: O(n), Space Complexity: O(1)
func (r *Reader) readWithBuffer(buf []byte, n int) int {
	totalRead := 0

	for totalRead < n {
		// If we have buffered characters, use them first
		if r.buf4Index < r.buf4Size {
			buf[totalRead] = r.buf4[r.buf4Index]
			totalRead++
			r.buf4Index++
			continue
		}

		// Buffer is empty, read more using read4
		r.buf4Size = read4(r.buf4)
		r.buf4Index = 0

		// If nothing was read, we're done
		if r.buf4Size == 0 {
			break
		}
	}

	return totalRead
}

// readIterative is a simple iterative approach
// Time Complexity: O(n), Space Complexity: O(1)
func readIterative(buf []byte, n int) int {
	totalRead := 0

	for totalRead < n {
		// Create a temporary buffer for read4
		tempBuf := make([]byte, 4)
		readCount := read4(tempBuf)

		if readCount == 0 {
			break
		}

		// Copy as much as we can
		for i := 0; i < readCount && totalRead < n; i++ {
			buf[totalRead] = tempBuf[i]
			totalRead++
		}
	}

	return totalRead
}

// readUsingPointer uses pointer arithmetic style approach
// Time Complexity: O(n), Space Complexity: O(1)
func readUsingPointer(buf []byte, n int) int {
	// Convert to pointer for manual copying
	bufPtr := 0
	tempBuf := make([]byte, 4)

	for bufPtr < n {
		readCount := read4(tempBuf)

		if readCount == 0 {
			break
		}

		// Copy characters
		for i := 0; i < readCount && bufPtr < n; i++ {
			buf[bufPtr] = tempBuf[i]
			bufPtr++
		}
	}

	return bufPtr
}