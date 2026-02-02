package tinyurl

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TinyURLService represents a URL shortening service
type TinyURLService struct {
	mu          sync.RWMutex
	urlToCode   map[string]string // Original URL -> short code
	codeToURL   map[string]string // Short code -> original URL
	codeToStats map[string]*URLStats // Short code -> usage statistics
	alphabet    string // Characters allowed in short codes
	codeLength  int    // Length of short codes
}

// URLStats represents usage statistics for a shortened URL
type URLStats struct {
	mu            sync.RWMutex
	CreatedAt     time.Time
	AccessCount   int64
	LastAccessed  time.Time
	UserAgents    map[string]int64 // User agent -> count
	Referrers     map[string]int64 // Referrer -> count
	Countries     map[string]int64 // Country code -> count
}

// NewTinyURLService creates a new URL shortening service
func NewTinyURLService() *TinyURLService {
	// Base62 alphabet: 0-9, A-Z, a-z
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	
	return &TinyURLService{
		urlToCode:   make(map[string]string),
		codeToURL:   make(map[string]string),
		codeToStats: make(map[string]*URLStats),
		alphabet:    alphabet,
		codeLength:  7, // 7 characters gives 62^7 ≈ 3.5 trillion possible codes
	}
}

// Shorten creates a short code for a given URL
func (s *TinyURLService) Shorten(url string) (string, error) {
	if url == "" {
		return "", errors.New("URL cannot be empty")
	}
	
	// Check if URL already has a short code
	s.mu.RLock()
	if code, exists := s.urlToCode[url]; exists {
		s.mu.RUnlock()
		return code, nil
	}
	s.mu.RUnlock()
	
	// Generate unique short code
	code := s.generateUniqueCode(url)
	
	// Store mapping
	s.mu.Lock()
	s.urlToCode[url] = code
	s.codeToURL[code] = url
	s.codeToStats[code] = &URLStats{
		CreatedAt:    time.Now(),
		AccessCount:  0,
		LastAccessed: time.Time{},
		UserAgents:   make(map[string]int64),
		Referrers:    make(map[string]int64),
		Countries:    make(map[string]int64),
	}
	s.mu.Unlock()
	
	return code, nil
}

// Expand retrieves the original URL for a given short code
func (s *TinyURLService) Expand(code string) (string, error) {
	if code == "" {
		return "", errors.New("code cannot be empty")
	}
	
	s.mu.RLock()
	url, exists := s.codeToURL[code]
	s.mu.RUnlock()
	
	if !exists {
		return "", errors.New("code not found")
	}
	
	return url, nil
}

// Redirect retrieves the original URL and records access statistics
func (s *TinyURLService) Redirect(code, userAgent, referrer, country string) (string, error) {
	if code == "" {
		return "", errors.New("code cannot be empty")
	}
	
	s.mu.RLock()
	url, exists := s.codeToURL[code]
	stats, statsExist := s.codeToStats[code]
	s.mu.RUnlock()
	
	if !exists {
		return "", errors.New("code not found")
	}
	
	// Update statistics
	if statsExist {
		stats.mu.Lock()
		stats.AccessCount++
		stats.LastAccessed = time.Now()
		
		if userAgent != "" {
			stats.UserAgents[userAgent]++
		}
		if referrer != "" {
			stats.Referrers[referrer]++
		}
		if country != "" {
			stats.Countries[country]++
		}
		stats.mu.Unlock()
	}
	
	return url, nil
}

// GetStats retrieves statistics for a short code
func (s *TinyURLService) GetStats(code string) (*URLStats, error) {
	if code == "" {
		return nil, errors.New("code cannot be empty")
	}
	
	s.mu.RLock()
	stats, exists := s.codeToStats[code]
	s.mu.RUnlock()
	
	if !exists {
		return nil, errors.New("code not found")
	}
	
	// Return a copy to avoid concurrent modification
	stats.mu.RLock()
	defer stats.mu.RUnlock()
	
	return &URLStats{
		CreatedAt:     stats.CreatedAt,
		AccessCount:   stats.AccessCount,
		LastAccessed:  stats.LastAccessed,
		UserAgents:    copyMap(stats.UserAgents),
		Referrers:     copyMap(stats.Referrers),
		Countries:     copyMap(stats.Countries),
	}, nil
}

// Delete removes a short code and its associated data
func (s *TinyURLService) Delete(code string) error {
	if code == "" {
		return errors.New("code cannot be empty")
	}
	
	s.mu.Lock()
	defer s.mu.Unlock()
	
	url, exists := s.codeToURL[code]
	if !exists {
		return errors.New("code not found")
	}
	
	delete(s.urlToCode, url)
	delete(s.codeToURL, code)
	delete(s.codeToStats, code)
	
	return nil
}

// GetTotalURLs returns the total number of shortened URLs
func (s *TinyURLService) GetTotalURLs() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urlToCode)
}

// generateUniqueCode generates a unique short code for a URL
func (s *TinyURLService) generateUniqueCode(url string) string {
	// Use MD5 hash of URL + timestamp for uniqueness
	hashInput := fmt.Sprintf("%s%d", url, time.Now().UnixNano())
	hash := md5.Sum([]byte(hashInput))
	hashStr := hex.EncodeToString(hash[:])
	
	// Convert hash to base62
	var code string
	for i := 0; i < s.codeLength; i++ {
		// Take bytes from hash and map to alphabet
		idx := int(hashStr[i%len(hashStr)]) % len(s.alphabet)
		code += string(s.alphabet[idx])
	}
	
	// Check for collisions (very rare but possible)
	s.mu.RLock()
	_, exists := s.codeToURL[code]
	s.mu.RUnlock()
	
	// If collision, append random characters
	if exists {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 3; i++ {
			idx := rand.Intn(len(s.alphabet))
			code += string(s.alphabet[idx])
		}
	}
	
	return code
}

// copyMap creates a copy of a string->int64 map
func copyMap(original map[string]int64) map[string]int64 {
	copy := make(map[string]int64)
	for k, v := range original {
		copy[k] = v
	}
	return copy
}

// Example usage:
// service := NewTinyURLService()
// code, err := service.Shorten("https://example.com/very/long/url")
// if err != nil { ... }
// url, err := service.Redirect(code, "Mozilla/5.0", "https://google.com", "US")
// if err != nil { ... }
// stats, err := service.GetStats(code)
// if err != nil { ... }