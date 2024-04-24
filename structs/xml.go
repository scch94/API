package structs

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Header  string   `xml:"Header"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	XMLName xml.Name `xml:"Body"`
	Send    Send     `xml:"send"`
}

type Send struct {
	XMLName       xml.Name `xml:"send"`
	Mobile        string   `xml:"mobile"`
	Message       string   `xml:"message"`
	UseOriginName string   `xml:"useOriginName"`
}
