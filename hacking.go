package main
import (
	"encoding/hex"
	"fmt"
	"sync"
	"crypto/sha1"
)
//PW IS MONSTER!!!!!
func main() {
	pwChan := make(chan bool,1)
	var sha [20]byte
	lowerCaseLetters := []rune("abcdefghijklmnopqrstuvwxyz")
	word := []rune("")
	pw := "0e818bfa0679df304036382aaa7667df92cbe30e"
	copy(sha[:], pw)
	go genWord(lowerCaseLetters,word,len(lowerCaseLetters),7,pw,pwChan)
	<-pwChan
}
func genWord(charSet []rune,word []rune,setSize int,wordSize int,pw string,pwChan chan bool){
	var wg sync.WaitGroup
	if(wordSize == 0){
		checkPw(word,pw,pwChan)
		return		
	}
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i := 0;i < setSize; i++{
			newWord := append(word,charSet[i])
			genWord(charSet,newWord,setSize,wordSize-1,pw,pwChan)	
		}
	}()
	wg.Wait()
}
func checkPw(word []rune,pw string,pwChan chan bool)bool{
	wordString := string(word)
	shaArray := sha1.Sum([]byte(wordString))
	shaString := hex.EncodeToString(shaArray[:])
	if shaString == pw {
		fmt.Println(string(word))
		pwChan <- true
		return true
	}
	return false
}