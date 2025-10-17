package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

// type Hangman struct {
// 	secretWord       string
// 	guessedLetters   []byte
// 	chancesRemaining uint
// 	correctGuesses   []byte
// }

// func NewGame(secretWord string) Hangman {
// 	return Hangman{
// 		secretWord:       secretWord,
// 		guessedLetters:   []byte{},
// 		chancesRemaining: 7,
// 		correctGuesses:   []byte{},
// 	}

// }
func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getSecretWord(wordFileName string) string {
	allowedwords := []string{}
	wordfile, err := os.Open(wordFileName)
	if err != nil {
		fmt.Println("The file could not open ", err)
	}
	defer wordfile.Close()
	scanner := bufio.NewScanner(wordfile)
	for scanner.Scan() {
		word := scanner.Text()
		if word == strings.ToLower(word) && IsLetter(word) && len(word) >= 6 {
			allowedwords = append(allowedwords, word)
		}

	}

	randomNum := rand.Intn(len(allowedwords))
	return allowedwords[randomNum]

}

// func checkguess(state Hangman, guessedLetter byte) Hangman {
// 	return state
// }

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))

}
