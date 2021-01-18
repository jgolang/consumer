package soap

import (
	"encoding/xml"
	"time"
)

// Version doc
type Version int

// RequestInfo doc
type RequestInfo struct {
	URL     string
	Action  string
	Header  interface{}
	Content interface{}
	Timeout time.Duration
}

// ResponseData doc
type ResponseData struct {
	Header  interface{}
	Content interface{}
}

// Envelope11 XML SOAP 1.1
type Envelope11 struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *Header
	Body    *Body
}

// Header Header of a SOAP message
type Header struct {
	XMLName xml.Name    `xml:"Header"`
	Content interface{} `xml:",any"`
}

// Body Body of a SOAP message
type Body struct {
	XMLName xml.Name    `xml:"Body"`
	Content interface{} `xml:",any"`
}

// Envelope12 XML SOAP 1.2
type Envelope12 struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header  *Header
	Body    *Body
}
