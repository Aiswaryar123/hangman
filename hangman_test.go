package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func createDictFile(words []string) (string, error) {
	f, err := os.CreateTemp("/tmp", "hangman-dict")
	if err != nil {
		fmt.Println("Couldn't create temp file.")
	}
	data := strings.Join(words, "\n")
	_, err = f.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}
func TestSecretWordNoCapitals(t *testing.T) {
	wordList, err := createDictFile([]string{"Lion", "Elephant", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}
func TestSecretWordNoPunc(t *testing.T) {
	wordList, err := createDictFile([]string{"Lion's", "Elephant's", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}
func TestSecretWordLength(t *testing.T) {
	wordList, err := createDictFile([]string{"lion", "pen", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}

}
func TestCorrectguess(t *testing.T) {

	userInput := 'p'
	currentstate := Hangman{
		secretWord:       "police",
		guessedLetters:   []byte{'a', 'c'},
		chancesRemaining: 5,
		correctGuesses:   []byte{},
	}
	newstate := checkguess(currentstate, byte(userInput))
	expected := Hangman{
		secretWord:       "police",
		guessedLetters:   append(currentstate.guessedLetters, byte(userInput)),
		chancesRemaining: 5,
		correctGuesses:   append((currentstate.correctGuesses), byte(userInput)),
	}

	if newstate.secretWord != expected.secretWord {
		t.Errorf("secret word was modified")
	}
	if string(newstate.guessedLetters) != string(expected.guessedLetters) {
		t.Errorf("Error processing guessed letters\n")
	}
	if string(newstate.correctGuesses) != string(expected.correctGuesses) {
		t.Errorf("Error processing Correctly guessed letters!\n")

	}
	if newstate.chancesRemaining != expected.chancesRemaining {
		t.Errorf("Chances left modified\n")
	}
}
func TestIncorrectguess(t *testing.T) {
	userInput := 'a'
	currentstate := Hangman{
		secretWord:       "police",
		guessedLetters:   []byte{},
		chancesRemaining: 7,
		correctGuesses:   []byte{},
	}
	newstate := checkguess(currentstate, byte(userInput))
	expected := Hangman{
		secretWord:       "police",
		guessedLetters:   append(currentstate.guessedLetters, byte(userInput)),
		chancesRemaining: 6,
		correctGuesses:   currentstate.correctGuesses,
	}
	if newstate.secretWord != expected.secretWord {
		t.Errorf("secret word was modified")
	}
	if string(newstate.guessedLetters) != string(expected.guessedLetters) {
		t.Errorf("Error processing guessed letters\n")
	}
	if string(newstate.correctGuesses) != string(expected.correctGuesses) {
		t.Errorf("Error processing Correctly guessed letters!\n")

	}
	if newstate.chancesRemaining != expected.chancesRemaining {
		t.Errorf("Chances left modified\n")
	}
}
