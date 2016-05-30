package randomNamesArg

import (
    "fmt"
    "testing"
    "time"
    "math/rand"
)


func TestFetchNyS(t *testing.T) {
    //Seedeamos al generador random   
   	rand.Seed(time.Now().UnixNano())
       
    for i := 0; i < 20; i++ {
       
        fmt.Println(FetchNyS())
    }
    
}