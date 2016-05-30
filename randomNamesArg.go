package randomNamesArg

import(
    "os"
    "bufio"
    "io/ioutil"
    "strings"
    "fmt"
    "math/rand"
)
//FetchNyS devuelve un nombre y apellido al azar
//asegurarse se seedear al random !!!
func FetchNyS()string  {
    
    var listaNombres,listaApellidos  []string
    var randomNombre,randomApellido  string

     
             
   //Abrimos los archivos
   fileNombres, err := os.Open("./listadoNombres.txt")
   fileApellidos, err := os.Open("./listadoApellidos.txt")
   if err != nil {
    fmt.Printf("error opening file: %v\n",err)
    os.Exit(1)
   }
   //Agarramos los readers
   readerNombres:= bufio.NewReader(fileNombres)
   readerApellidos:= bufio.NewReader(fileApellidos)    
   
   //Leemos todo el archivos y los partimos
   textNombres,err:=ioutil.ReadAll(readerNombres)
   textApellidos,err:=ioutil.ReadAll(readerApellidos)
   if err != nil {
    fmt.Printf("error reading: %v\n",err)
    os.Exit(1)
   }
   
   //Si no le podes el char "0x" (CR), tu vida va ser una pesadilla en el sprintf
   listaNombres=strings.Split(string(textNombres),"\x0d\n")
   listaApellidos=strings.Split(string(textApellidos),"\x0d\n")   
   
   //sacamos los random
   randomNombre=listaNombres[rand.Intn(len(listaNombres))]
   randomApellido=listaApellidos[rand.Intn(len(listaApellidos))]

   return fmt.Sprintf("%v %v",randomNombre,randomApellido) 
  
}
