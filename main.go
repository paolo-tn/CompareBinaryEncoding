package main

import (
  "fmt"
  "encoding/json"
  "encoding/hex"
  "os"
  "log"
  "github.com/vmihailenco/msgpack/v5"
 )

type User struct {
  UserName string     `json:"userName"`
  FavoriteNumber int  `json:"favoriteNumber"`
  Interests []string `json:"interests"`
}
func main(){
  
  fmt.Printf("%s", "Start function main")
  
  fmt.Println("")
  
  u:= User{
    UserName : "Martin",
    FavoriteNumber: 1337,
    Interests : []string{"daydreaming","hacking"},
  }
  
  //fmt.Printf("%+v",u)
  
  enc := json.NewEncoder(os.Stdout)
  if err := enc.Encode(u); err!=nil{      
    //fmt.Printf("%x",b)
    log.Fatal(err)
  }
  
  b,err := json.Marshal(u)
  
  if err!= nil{
    
  }
  fmt.Println(">>>>> Print array of bytes binary representation <<<<<<<")
  fmt.Printf("%08b",b)
  //os.Stdout.Write(b)
  fmt.Println("")
  fmt.Println("")
  fmt.Println(">>>>> Print array of bytes as hex representation <<<<<<<")
  fmt.Printf("Array of bytes (JSON) is %d long:\n\n%x",len(b),b)
  fmt.Println("")
  fmt.Println("")
  encodedString := hex.EncodeToString(b)
  fmt.Printf("Array of bytes (JSON) hex representation %s", encodedString )
  fmt.Println("")
  fmt.Println("")
  fmt.Println(">>>>> MessagePack <<<<<<<<<<<<<<<")
  bMsg, errMsg := msgpack.Marshal(u)
  if(errMsg!= nil){
    log.Fatal(errMsg)
  }
  fmt.Printf("Array of bytes (MessagePack) is %d long:\n\n%x",len(bMsg),bMsg)

}


