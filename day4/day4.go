package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func ParseRow(row string) []string {
	return strings.Fields(row)
}

func ContainsDuplicates(passphrase []string) bool {
	for i, word1 := range passphrase {
		for j, word2 := range passphrase {
			if i == j {
				continue
			}

			if word1 == word2 {
				return true
			}
		}
	}
	return false
}

func ContainsAnagrams(passphrase []string) bool {
	for i, word1 := range passphrase {
		for j, word2 := range passphrase {
			if i == j {
				continue
			}

			if IsAnagram(word1, word2) {
				return true
			}
		}
	}
	return false
}

func IsPassphraseValid(row string) bool {
	return !ContainsDuplicates(ParseRow(row))
}

func IsPassphraseValidWithoutAnagrams(row string) bool {
	parsed := ParseRow(row)
	return !ContainsDuplicates(parsed) && !ContainsAnagrams(parsed)
}

func IsAnagram(first, second string) bool {
	hashFirst := make(map[rune]int)
	for _, c := range first {
		count, ok := hashFirst[c]
		if !ok {
			hashFirst[c] = 1
		} else {
			hashFirst[c] = count + 1
		}
	}

	hashSecond := make(map[rune]int)
	for _, c := range second {
		count, ok := hashSecond[c]
		if !ok {
			hashSecond[c] = 1
		} else {
			hashSecond[c] = count + 1
		}
	}

	if len(hashFirst) != len(hashSecond) {
		return false
	}

	for keyA, valueA := range hashFirst {
		valueB, ok := hashSecond[keyA]

		if !ok || valueA != valueB {
			return false
		}
	}

	return true
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	validCount := 0
	validWithoutAnagrams := 0

	for scanner.Scan() {
		row := scanner.Text()

		if IsPassphraseValid(row) {
			validCount += 1
		}

		if IsPassphraseValidWithoutAnagrams(row) {
			validWithoutAnagrams += 1
		}
	}

	fmt.Println("Valid passphrases", validCount)
	fmt.Println("Valid without anagrams", validWithoutAnagrams)
}
