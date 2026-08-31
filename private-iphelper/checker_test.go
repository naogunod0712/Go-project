package iphelper

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type resolverStub struct {
	addresses []string
}

func (resolver resolverStub) LookupHost(context.Context, string) ([]string, error) {
	return resolver.addresses, nil
}

func TestCheckerReportsHealthyWebsite(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	checker := &Checker{
		HTTPClient: server.Client(),
		Resolver:   resolverStub{addresses: []string{"192.0.2.10"}},
	}
	result, err := checker.Check(context.Background(), server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if !result.Healthy || result.StatusCode != http.StatusOK {
		t.Fatalf("unexpected health result: %+v", result)
	}
	if len(result.ResolvedAddresses) != 1 || result.ResolvedAddresses[0] != "192.0.2.10" {
		t.Fatalf("unexpected DNS result: %+v", result.ResolvedAddresses)
	}
}

func TestCheckerReportsUnhealthyStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	checker := &Checker{
		HTTPClient: server.Client(),
		Resolver:   resolverStub{addresses: []string{"192.0.2.10"}},
	}
	result, err := checker.Check(context.Background(), server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if result.Healthy || result.StatusCode != http.StatusInternalServerError {
		t.Fatalf("unexpected health result: %+v", result)
	}
}

func TestCheckerRejectsInvalidURL(t *testing.T) {
	_, err := NewChecker().Check(context.Background(), "not-a-url")
	if err == nil {
		t.Fatal("expected invalid URL error")
	}
}
