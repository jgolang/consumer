package rest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jgolang/consumer"
	// "github.com/jgolang/consumer"
)

// ConsumeRestService doc..
func ConsumeRestService(requestInfo RequestInfo, v interface{}) (*http.Response, error) {
	consumer.Print("URL: ", requestInfo.Endpoint)
	consumer.Print("HEADERS: %v", requestInfo.Headers)
	var buff []byte
	var err error
	if requestInfo.Body != nil {
		buff, err = json.Marshal(requestInfo.Body)
		if err != nil {
			return nil, err
		}
		consumer.LogRequest(buff)
	}
	// else {
	// 	buff = make([]uint8, 0)
	// }
	request, err := http.NewRequest(requestInfo.Method, requestInfo.Endpoint, bytes.NewReader(buff))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	for key, value := range requestInfo.Headers {
		request.Header.Add(key, value)
	}
	queryParameters := request.URL.Query()
	for key, value := range requestInfo.Query {
		queryParameters.Add(key, value)
	}
	request.URL.RawQuery = queryParameters.Encode()
	consumer.Print("QUERY STRING: ", request.RequestURI)
	timeout := requestInfo.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	responseBuff, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if responseBuff != nil && len(responseBuff) != 0 {
		consumer.LogResponse(responseBuff)
		err = json.Unmarshal(responseBuff, &v)
		if err != nil {
			return nil, err
		}
	}
	return response, nil
}
