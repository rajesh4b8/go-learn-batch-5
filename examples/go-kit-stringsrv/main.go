package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json: "v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type StringService interface {
	Uppercase(string) (string, error)
	// Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty string")
	}

	return strings.ToUpper(s), nil
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}

		return UppercaseResponse{strings.ToUpper(req.S), ""}, nil
	}
}

func decoderUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func main() {
	svc := stringService{}
	uppercaseHandler := httptransport.NewServer(makeUppercaseEndpoint(svc), decoderUppercaseRequest, encodeResponse)
	http.Handle("/uppercase", uppercaseHandler)
	// http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
