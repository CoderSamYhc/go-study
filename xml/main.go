package main

import (
	"encoding/xml"
	"fmt"
	"go-study/xml/book"
	"io/ioutil"
	"os"
)


func encodeXml() {
	file, err := os.Open("decode.xml")
	if err != nil {
		fmt.Println("file err:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read err:", err)
		return
	}

	v := book.EncodeBookStore{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("unmarshal err:", err)
		return
	}
	fmt.Printf("%v", v.Test[0].ItemName)
}

func decodeXml() {
	bs := &book.DecodeBookStore{Version: "1"}
	bs.Store = append(bs.Store, book.DecodeBook{"golang实战", "$450"})
	bs.Store = append(bs.Store, book.DecodeBook{"三体", "$150"})
	bs.Test = append(bs.Test, book.DecodeTest{"123", "456"})
	bs.Test = append(bs.Test, book.DecodeTest{"31", "426"})

	output, err := xml.MarshalIndent(bs, " ", "  ")
	if err != nil {
		fmt.Println("marshal err:", err)
		return
	}

	file, err := os.OpenFile("decode.xml", os.O_APPEND | os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("file err:", err)
		return
	}
	defer file.Close()

	file.Write([]byte(xml.Header))
	file.Write(output)
}

func main() {

	encodeXml()
	//decodeXml()
}
