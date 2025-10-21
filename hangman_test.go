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
		chancesRemaining: 6,
		correctGuesses:   []byte{'c'},
	}
	newstate := checkguess(currentstate, byte(userInput))
	expected := Hangman{
		secretWord:       "police",
		guessedLetters:   append(currentstate.guessedLetters, byte(userInput)),
		chancesRemaining: 6,
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
func TestIncorrectguess1(t *testing.T) {
	userInput := 'a'
	currentstate := Hangman{
		secretWord:       "police",
		guessedLetters:   []byte{'x', 'y', 'z'},
		chancesRemaining: 4,
		correctGuesses:   []byte{},
	}
	newstate := checkguess(currentstate, byte(userInput))
	expected := Hangman{
		secretWord:       "police",
		guessedLetters:   append(currentstate.guessedLetters, byte(userInput)),
		chancesRemaining: 3,
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

func TestAlreadyguess(t *testing.T) {

	userInput := 'a'
	currentstate := Hangman{
		secretWord:       "police",
		guessedLetters:   []byte{'a', 'c'},
		chancesRemaining: 6,
		correctGuesses:   []byte{'c'},
	}
	newstate := checkguess(currentstate, byte(userInput))
	expected := Hangman{
		secretWord:       currentstate.secretWord,
		guessedLetters:   currentstate.guessedLetters,
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
func TestCheckWon(t *testing.T) {
	state := Hangman{
		secretWord:       "police",
		correctGuesses:   []byte{'p', 'o', 'l', 'i', 'c', 'e'},
		guessedLetters:   []byte{'p', 'o', 'l', 'i', 'c', 'e'},
		chancesRemaining: 7,
	}

	if !CheckWon(state) {
		t.Errorf("Expected true, but got false")
	}

	state2 := Hangman{
		secretWord:       "police",
		correctGuesses:   []byte{'p', 'o', 'l', 'c'},
		guessedLetters:   []byte{'p', 'o', 'l', 'c', 'a'},
		chancesRemaining: 6,
	}

	if CheckWon(state2) {
		t.Errorf("Expected false, but got true")
	}
}

func TestDisplay_allcorrect(t *testing.T) {
	state1 := Hangman{
		secretWord:       "police",
		correctGuesses:   []byte{'p', 'o', 'c'},
		guessedLetters:   []byte{'p', 'o', 'c', 'x'},
		chancesRemaining: 6,
	}
	expected1 := "po--c-"
	result1 := displayWord(state1)
	if result1 != expected1 {
		t.Errorf("Expected %s, got %s", expected1, result1)
	}
}
func TestDisplayWord_Somerepeatletters(t *testing.T) {
	state := Hangman{
		secretWord:       "letter",
		correctGuesses:   []byte{'e', 't'},
		guessedLetters:   []byte{'e', 't', 'x'},
		chancesRemaining: 6,
	}
	expected := "-ette-"
	result := displayWord(state)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
func TestDisplayWord_Nonecorrect(t *testing.T) {
	state := Hangman{
		secretWord:       "banana",
		correctGuesses:   []byte{},
		guessedLetters:   []byte{'x', 'y', 'z'},
		chancesRemaining: 4,
	}
	expected := "------"
	result := displayWord(state)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
func TestCheckLoss(t *testing.T) {
	state := Hangman{
		secretWord:       "police",
		correctGuesses:   []byte{'p', 'o', 'l'},
		guessedLetters:   []byte{'p', 'o', 'l', 'x', 'y', 'z'},
		chancesRemaining: 0,
	}

	if !CheckLoss(state) {
		t.Errorf("Expected true (loss), but got false")
	}

	state2 := Hangman{
		secretWord:       "police",
		correctGuesses:   []byte{'p', 'o'},
		guessedLetters:   []byte{'p', 'o', 'x'},
		chancesRemaining: 6,
	}

	if CheckLoss(state2) {
		t.Errorf("Expected false (not lost yet), but got true")
	}
}
