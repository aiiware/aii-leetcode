package arrays

// ValidPhoneNumbers solves LeetCode problem 0193: Valid Phone Numbers
// Difficulty: Easy
// Tags: String, Text Processing
//
// Given a text file file.txt that contains list of phone numbers (one per line),
// write a one liner bash script to print all valid phone numbers.
//
// A valid phone number must appear in one of the following two formats:
// (xxx) xxx-xxxx or xxx-xxx-xxxx
//
// Time complexity: O(n), Space complexity: O(1)
func ValidPhoneNumbers(phones []string) []string {
	var result []string
	
	for _, phone := range phones {
		if isValidPhoneNumber(phone) {
			result = append(result, phone)
		}
	}
	
	return result
}

func isValidPhoneNumber(phone string) bool {
	// Format: (xxx) xxx-xxxx (length 14)
	// (123) 456-7890
	// 01234567890123
	if len(phone) == 14 {
		if phone[0] == '(' && phone[4] == ')' && phone[5] == ' ' && phone[9] == '-' {
			// Check digits
			if isDigit(phone[1]) && isDigit(phone[2]) && isDigit(phone[3]) &&
				isDigit(phone[6]) && isDigit(phone[7]) && isDigit(phone[8]) &&
				isDigit(phone[10]) && isDigit(phone[11]) && isDigit(phone[12]) && isDigit(phone[13]) {
				return true
			}
		}
	}
	
	// Format: xxx-xxx-xxxx (length 12)
	// 123-456-7890
	// 012345678901
	if len(phone) == 12 {
		if phone[3] == '-' && phone[7] == '-' {
			// Check digits
			if isDigit(phone[0]) && isDigit(phone[1]) && isDigit(phone[2]) &&
				isDigit(phone[4]) && isDigit(phone[5]) && isDigit(phone[6]) &&
				isDigit(phone[8]) && isDigit(phone[9]) && isDigit(phone[10]) && isDigit(phone[11]) {
				return true
			}
		}
	}
	
	return false
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
