package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"github.com/atotto/clipboard"
)

func main() {
	// Load the wordlist
	words, err := LoadWordlist("/Users/cfinlay/projects/golang/wordpassgen/cleaned_words.txt")
	if err != nil {
		log.Fatal(err)
	}

	password := GeneratePassword(words, 3) // 3 words
	fmt.Println("Generated password:", password)

	// Copy to clipboard
	err = clipboard.WriteAll(password)
	if err != nil {
		log.Fatalf("Failed to copy to clipboard: %v", err)
	}
	fmt.Println("Password copied to clipboard.")
}

// LoadWordlist reads the wordlist from a file and returns a slice of words
func LoadWordlist(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	var words []string
	for _, line := range lines {
		word := strings.TrimSpace(line)
		if word != "" {
			words = append(words, word)
		}
	}
	return words, nil
}

// GeneratePassword creates a password from N words with random number/symbol placement
func GeneratePassword(wordlist []string, wordCount int) string {
	rand.Seed(time.Now().UnixNano())

	if len(wordlist) < wordCount {
		log.Fatalf("Wordlist must contain at least %d words", wordCount)
	}

	// Select N random words
	words := make([]string, wordCount)
	for i := 0; i < wordCount; i++ {
		words[i] = wordlist[rand.Intn(len(wordlist))]
	}

	// Capitalize one random word
	capIndex := rand.Intn(wordCount)
	words[capIndex] = Capitalize(words[capIndex])

	// Join with hyphens
	password := strings.Join(words, "-")

	// Random number (2-3 digits)
	number := fmt.Sprintf("%d", rand.Intn(900)+100) // 100-999

	// Random symbol
	symbols := "!@#$%&*?"
	symbol := string(symbols[rand.Intn(len(symbols))])

	// Random insert positions
	positions := []int{rand.Intn(len(password) + 1), rand.Intn(len(password) + 2)}

	// Insert number and symbol
	password = InsertAt(password, number, positions[0])
	password = InsertAt(password, symbol, positions[1])

	return password
}

// Capitalize makes the first letter uppercase
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + word[1:]
}

// InsertAt inserts a string into another at a given index
func InsertAt(str, insert string, pos int) string {
	if pos > len(str) {
		pos = len(str)
	}
	return str[:pos] + insert + str[pos:]
}
