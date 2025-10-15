package main


import (
	"strings"
	"testing"
	"unicode"
)


func TestSecretWordNoCapitals(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if secretWord != strings.ToLower(secretWord) {
		t.Errorf("Should not get words with capital letters. Got %s", secretWord)
	}


}
func TestSecretWordNoPunctuation(t *testing.T){

	 wordlist := "/usr/share/dict/words"
	 secretword := getSecretWord(wordlist)

	for _, c := range secretword {
		if !unicode.IsLetter(c) {
			t.Errorf("Expected only letters but Got:%s",secretword)
			break
		}
	}
}


