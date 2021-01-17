package rest

import (
	"net/http"
	"reflect"
	"testing"
)

func TestConsumeRestService(t *testing.T) {
	query := make(map[string]string)
	query["limit"] = "20"
	query["term"] = "Nirvana"
	query["entity"] = "movie"
	query["country"] = "GT"
	type args struct {
		requestInfo RequestInfo
		v           interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			name: "Search itunes test",
			args: args{
				requestInfo: RequestInfo{
					Method:      http.MethodGet,
					Endpoint:    "https://itunes.apple.com/search",
					QueryParams: query,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConsumeRestService(tt.args.requestInfo, tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConsumeRestService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConsumeRestService() = %v, want %v", got, tt.want)
			}
		})
	}
}
