package main

import (
	"fmt"

	pb "github.com/skalle/helloproto/proto/hellopb"
)

func main() {
	fmt.Println("Hello world!")
	pb := &pb.HelloRequest{
		Name: "Opel Ascona",
	}

	fmt.Printf("Hello: %v\n", pb.Name)
	fmt.Println("Hello: Le World")

}
