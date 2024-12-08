package main

import (
  "fmt"
  "encoding/json"
  //"encoding/hex"
  "os"
  "log"
  "github.com/vmihailenco/msgpack/v5"
  "github.com/hamba/avro/v2"
 )

type User struct {
  UserName string     `json:"userName"`
  FavoriteNumber int  `json:"favoriteNumber"`
  Interests []string `json:"interests"`
}

type UserA struct {
  A string  `avro:"a"`
  B int  `avro:"b"` 
  C []string `avro:"c"`
}


func main(){
  
  fmt.Println("%s", "This is a simple demo of how serialization works with different protocols")
  fmt.Println("It was inspired by book \"Designing Data Intensive Application by Martin Kleppmann - Chapter 4: Encoding and Evolution\"")  
  fmt.Println("")
  
  //create an instance of the User object
  u:= User{
    UserName : "Martin",
    FavoriteNumber: 1337,
    Interests : []string{"daydreaming","hacking"},
  }
  
  fmt.Printf("This is the oject we want to serialize: %+v",u)
  fmt.Println("")
  fmt.Println("1. Serialize object as JSON String:")
  enc := json.NewEncoder(os.Stdout)
    //fmt.Printf("%x",b)
  if err := enc.Encode(u); err!=nil{      
    fmt.Println("error while marshalling with json")
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
  fmt.Printf("Byte array (JSON String) is %d long:\n\n%x",len(b),b)
  //alternative way to show the byte array:
  //encodedString := hex.EncodeToString(b)
  //fmt.Printf("Array of bytes (JSON) hex representation %s", encodedString )
  fmt.Println("")
  fmt.Println("")
  
  fmt.Println("")
  fmt.Println("")
  fmt.Println("2. Serialize object with MessagePack")
  bMsg, errMsg := msgpack.Marshal(u)
  if(errMsg!= nil){
    fmt.Println("error while marshalling with MessagePack")
    log.Fatal(errMsg)
  }
  fmt.Printf("Byte array (MessagePack) is %d long:\n\n%x",len(bMsg),bMsg)

  fmt.Println("")
  fmt.Println("")
  
  fmt.Println("3. Serialize object with Avro")
  
   
  schema, err := avro.Parse(`{
		"type": "record",
		"name": "simple",
		"namespace": "org.hamba.avro",
		"fields" : [
			{"name": "a", "type": "string"},
			{"name": "b", "type": "int"},
      {"name": "c", "type": {"type": "array","items": "string"}}
		]
	}`)
  
	if err != nil {
    fmt.Println("error:")
		log.Fatal(err)
	}
  fmt.Printf("Avro schema for the object is\n: %+v\n", schema)
  
  //object instance as per AVRO schema
  uA := UserA{A: "Martin",B: 1337,C:[]string{"daydreaming","hacking"}}
  
  fmt.Println("")
  
  bAvro, err := avro.Marshal(schema, uA)
  if err != nil{
    fmt.Println("error while marshalling with avro")
		log.Fatal(err)
  }
  fmt.Printf("Byte Array (AVRO) is %d long:\n\n%x",len(bAvro),bAvro)
 
}


