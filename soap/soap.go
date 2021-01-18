package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/jgolang/consumer"
)

// Soap11 version
const Soap11 Version = 11

// Soap12 version
const Soap12 Version = 12

// ConsumeSOAP11Service consumes a SOAP 1.2 web services
// expr: xmlquery expresion to find xml node
func ConsumeSOAP11Service(soapRequest RequestInfo, expr string) (*xmlquery.Node, error) {
	requestEnvelope := Envelope11{
		Header: &Header{
			Content: soapRequest.Header,
		},
		Body: &Body{
			Content: soapRequest.Content,
		},
	}
	xmlNode, _, err := ConsumeSOAPService(Soap11, soapRequest, &requestEnvelope)
	if err != nil {
		return nil, err
	}
	if expr == "" {
		return xmlNode, nil
	}
	node := xmlquery.FindOne(xmlNode, expr)
	if node == nil {
		return nil, fmt.Errorf("Not result")
	}
	consumer.Print("[SOAP integration] XML NODE: %v", node.InnerText())
	return node, nil
}

// ConsumeSOAP12Service consumes a SOAP 1.2 web services
// expr: xmlquery expresion to find xml node
func ConsumeSOAP12Service(soapRequest RequestInfo, expr string) (*xmlquery.Node, error) {
	requestEnvelope := Envelope12{
		Header: &Header{
			Content: soapRequest.Header,
		},
		Body: &Body{
			Content: soapRequest.Content,
		},
	}
	xmlNode, _, err := ConsumeSOAPService(Soap12, soapRequest, &requestEnvelope)
	if err != nil {
		return nil, err
	}
	if expr == "" {
		return xmlNode, nil
	}
	node := xmlquery.FindOne(xmlNode, expr)
	if node == nil {
		return nil, fmt.Errorf("Not result")
	}
	consumer.Print("[SOAP integration] XML NODE: %v", node.InnerText())
	return node, nil
}

// ConsumeSOAPService return *xmlquery.Node to find xml nodes in document
func ConsumeSOAPService(version Version, requestInfo RequestInfo, requestEnvelope Envelope) (*xmlquery.Node, *http.Response, error) {
	consumer.Print("[SOAP integration] ACTION: %v", requestInfo.Action)
	buff, err := xml.Marshal(requestEnvelope)
	if err != nil {
		return nil, nil, err
	}
	consumer.LogRequest(buff)
	reader := bytes.NewReader(buff)
	httpRequest, err := http.NewRequest("POST", requestInfo.URL, reader)
	if err != nil {
		return nil, nil, err
	}
	switch version {
	case Soap11:
		httpRequest.Header.Add("Content-Type", "text/xml; charset=utf-8")
		httpRequest.Header.Add("SOAPAction", requestInfo.Action)
	case Soap12:
		httpRequest.Header.Add("Content-Type", "application/soap+xml;charset=UTF-8")
	default:
		return nil, nil, fmt.Errorf("Error: no valid soap version. Invalid: %v", version)
	}
	timeout := requestInfo.Timeout
	if timeout == 0 {
		timeout = consumer.DefaultTimeout
	}
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Do(httpRequest)
	if err != nil {
		return nil, nil, err
	}
	responseBodyBuff, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	var xmlNode *xmlquery.Node
	if responseBodyBuff != nil && len(responseBodyBuff) != 0 {
		consumer.LogResponse(responseBodyBuff)
		xmlNode, err = xmlquery.Parse(strings.NewReader(string(responseBodyBuff)))
		if err != nil {
			return nil, nil, err
		}

	}
	return xmlNode, response, nil
}
