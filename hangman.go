package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

type Hangman struct {
	secretWord       string
	guessedLetters   []byte
	chancesRemaining uint
	correctGuesses   []byte
}

func NewGame(secretWord string) Hangman {
	return Hangman{
		secretWord:       secretWord,
		guessedLetters:   []byte{},
		chancesRemaining: 7,
		correctGuesses:   []byte{},
	}

}
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
func checkguess(state Hangman, guessedLetter byte) Hangman {

	isAlreadyGuessed := false
	for _, l := range state.guessedLetters {
		if l == guessedLetter {
			isAlreadyGuessed = true
			break
		}
	}
	isContainByte := strings.ContainsRune(state.secretWord, rune(guessedLetter))

	if isContainByte && !isAlreadyGuessed && state.chancesRemaining > 0 {
		state = Hangman{
			secretWord:       state.secretWord,
			guessedLetters:   append(state.guessedLetters, guessedLetter),
			correctGuesses:   append(state.correctGuesses, guessedLetter),
			chancesRemaining: state.chancesRemaining,
		}
		return state
	} else if !isContainByte && !isAlreadyGuessed && state.chancesRemaining > 0 {

		state = Hangman{
			secretWord:       state.secretWord,
			guessedLetters:   append(state.guessedLetters, guessedLetter),
			correctGuesses:   state.correctGuesses,
			chancesRemaining: state.chancesRemaining - 1,
		}
		return state
	} else {

		return state
	}
}
func CheckWon(state Hangman) bool {
	for _, ch := range state.secretWord {
		if !strings.ContainsRune(string(state.correctGuesses), ch) {
			return false
		}
	}
	return true
}
func displayWord(state Hangman) string {
	display := ""

	for _, ch := range state.secretWord {

		found := false
		for _, guessed := range state.correctGuesses {
			if guessed == byte(ch) {
				found = true
				break
			}
		}

		if found {
			display += string(ch)
		} else {
			display += "-"
		}
	}
	return display
}
