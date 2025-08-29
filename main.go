package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
	"example.com/protobuf-demo/pb"

)

func main() {
	p := &pb.Person(Id: 101, Name: "Vish", Email: "vish@example.com")

	bin,err := proto.Marshal(p)
	if err != nil {
		log.Fatal("marshal:", err)
	}
	fmt.Println("Encoded bytes:", len(bin))

	var out pb.Person
	if err := proto.Unmarshal(bin, &out); err != nil {
		log.Fatal("unmarshal:", err)
	}
	fmt.Printf("Decoded: id=%d name=%q email=%q\n",out.Id,out.Name, out.Email)
}