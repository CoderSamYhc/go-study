package book

import "encoding/xml"

type encodeBook struct {
	XMLName   xml.Name `xml:"book"`
	BookName  string   `xml:"bookName"`
	BookPrice string   `xml:"bookPrice"`
}

type encodeTest struct {
	XMLName  xml.Name `xml:"test"`
	ItemName string   `xml:"itemName"`
	ItemArgs string   `xml:"itemArgs"`
}
type EncodeBookStore struct {
	XMLName     xml.Name     `xml:"books"`
	Version     string       `xml:"version,attr"`
	Store       []encodeBook `xml:"book"`
	Test        []encodeTest `xml:"test"`
	Description string       `xml:",innerxml"`
}

type DecodeBook struct {
	BookName  string   `xml:"bookName"`
	BookPrice string   `xml:"bookPrice"`
}

type DecodeTest struct {
	ItemName string   `xml:"itemName"`
	ItemArgs string   `xml:"itemArgs"`
}
type DecodeBookStore struct {
	XMLName     xml.Name     `xml:"books"`
	Version     string       `xml:"version,attr"`
	Store       []DecodeBook `xml:"book"`
	Test        []DecodeTest `xml:"test"`
}

