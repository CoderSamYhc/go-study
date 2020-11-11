package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-study/protoes/protoes"
)

func main() {
	helloRequest := protoes.HelloRequest{
		Name: *proto.String("Sam"),
		UCount: *proto.Int32(21),
	}

	data, err := proto.Marshal(&helloRequest)

	if err != nil {
		fmt.Println("marshal err", err)
		return
	}

	fmt.Println("data", data)

	var list protoes.HelloRequest
	err = proto.Unmarshal(data, &list)

	if err != nil {
		fmt.Println("unmarshal err", err)
		return
	}
	fmt.Println("name", list.GetName())
	fmt.Println("count", list.GetUCount())
}
