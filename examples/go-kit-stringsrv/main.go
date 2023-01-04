package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type StringRequest struct {
	S string `json:"s"`
}

type StringResponse struct {
	V   string `json: "v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty string")
	}

	return strings.ToUpper(s), nil
}
func (stringService) Count(s string) int {
	return len(s)
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(StringRequest)
		v := svc.Count(req.S)
		return StringResponse{strconv.Itoa(v), ""}, nil
	}
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(StringRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return StringResponse{v, err.Error()}, nil
		}

		return StringResponse{strings.ToUpper(req.S), ""}, nil
	}
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request StringRequest
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
	uppercaseHandler := httptransport.NewServer(makeUppercaseEndpoint(svc), decodeRequest, encodeResponse)
	countHandler := httptransport.NewServer(makeCountEndpoint(svc), decodeRequest, encodeResponse)
	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
