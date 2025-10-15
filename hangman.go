package main
 import ("fmt"
    //    "bufio"
	//    "os"
	)
         

 func getSecretWord(wordFileName string)string{

	// wordfile,err := os.Open(wordFileName)
	// if err !=nil{
	// 	fmt.Println("The file could not open ",err)
	// }
    // // defer wordfile.Close()
    // scanner := bufio.NewScanner(wordfile)
	// for scanner.Scan() {
    //     fmt.Println(scanner.Text())
	// }
    // randomnum := len()


      
	return "elephant"

 }
 func main(){
	fmt.Println(getSecretWord("/usr/share/dict/words"))


 }