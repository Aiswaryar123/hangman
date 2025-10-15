package main
 import ("fmt"
       "bufio"
	   "os"
	   "strings"
	   "math/rand"
	)
         
//  func Isvalid(s string)bool
      
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
		if word == strings.ToLower(word){
			allowedwords= append(allowedwords,word)
		}
        
 

	}
	
    randomNum := rand.Intn(len(allowedwords))
	return allowedwords[randomNum]

}


func main(){
	fmt.Println(getSecretWord("/usr/share/dict/words"))


 }