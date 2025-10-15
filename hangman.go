package main
 import ("fmt"
       "bufio"
	   "os"
	   "strings"
	   "math/rand"
	   "unicode"
	)
         
  func IsLetter(s string)bool{
	for _,r :=range s{
		if !unicode.IsLetter(r){
			return false
		}
	}
	return true
  }
      
 func getSecretWord(wordFileName string)string{
   allowedwords :=[]string{}
	wordfile,err := os.Open(wordFileName)
	if err !=nil{
		fmt.Println("The file could not open ",err)
	}
    defer wordfile.Close()
    scanner := bufio.NewScanner(wordfile)
	for scanner.Scan() {
		word := scanner.Text()
		if word == strings.ToLower(word)&& IsLetter(word){
			allowedwords= append(allowedwords,word)
		}
        
 

	}
	
    randomNum := rand.Intn(len(allowedwords))
	return allowedwords[randomNum]

}


func main(){
	fmt.Println(getSecretWord("/usr/share/dict/words"))


 }