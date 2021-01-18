package soap

import (
	"encoding/xml"
	"testing"
)

func TestConsumeSOAP12Service(t *testing.T) {
	type GetByName struct {
		XMLName xml.Name `xml:"tem:GetByName"`
		Tem     string   `xml:"xmlns:tem,attr"`
		Name    string   `xml:"tem:name"`
	}
	type args struct {
		soapRequest RequestInfo
		expr        string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test soap 11",
			args: args{
				soapRequest: RequestInfo{
					URL:    "http://www.crcind.com/csp/samples/SOAP.Demo.CLS",
					Action: "GetByNameResponse",
					Content: GetByName{
						Name: "john",
						Tem:  "http://tempuri.org",
					},
				},
				expr: "//SQL",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := ConsumeSOAP12Service(tt.args.soapRequest, tt.args.expr); (err != nil) != tt.wantErr {
				t.Errorf("ConsumeSOAP11Service() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
