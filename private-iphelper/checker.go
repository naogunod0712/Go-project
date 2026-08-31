// Package iphelper provides website health checks.
package iphelper

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DNSResolver is the part of net.Resolver used by Checker.
type DNSResolver interface {
	LookupHost(ctx context.Context, host string) ([]string, error)
}

// Result describes the checks performed against a website.
type Result struct {
	URL               string
	Healthy           bool
	StatusCode        int
	ResponseTime      time.Duration
	ResolvedAddresses []string
	Redirects         []string
	TLSValid          bool
}

// Checker performs DNS and HTTP(S) health checks.
type Checker struct {
	HTTPClient *http.Client
	Resolver   DNSResolver
}

// NewChecker returns a checker with a ten-second request timeout.
func NewChecker() *Checker {
	return &Checker{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Resolver:   net.DefaultResolver,
	}
}

// Check resolves and requests rawURL, returning a health report.
func (c *Checker) Check(ctx context.Context, rawURL string) (Result, error) {
	parsedURL, err := parseURL(rawURL)
	if err != nil {
		return Result{URL: rawURL}, err
	}

	addresses, err := c.resolver().LookupHost(ctx, parsedURL.Hostname())
	result := Result{URL: parsedURL.String(), ResolvedAddresses: addresses}
	if err != nil {
		return result, fmt.Errorf("DNS lookup for %q: %w", parsedURL.Hostname(), err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		return result, fmt.Errorf("create website request: %w", err)
	}

	client := c.httpClient()
	redirects := []string{}
	clientCopy := *client
	clientCopy.CheckRedirect = func(request *http.Request, previous []*http.Request) error {
		redirects = append(redirects, request.URL.String())
		return nil
	}

	started := time.Now()
	response, err := clientCopy.Do(request)
	result.ResponseTime = time.Since(started)
	result.Redirects = redirects
	if err != nil {
		return result, fmt.Errorf("request website: %w", err)
	}
	defer response.Body.Close()

	result.StatusCode = response.StatusCode
	result.TLSValid = parsedURL.Scheme != "https" || response.TLS != nil
	result.Healthy = response.StatusCode >= http.StatusOK && response.StatusCode < http.StatusBadRequest && result.TLSValid
	return result, nil
}

func parseURL(rawURL string) (*url.URL, error) {
	parsedURL, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil || parsedURL.Host == "" || parsedURL.Hostname() == "" {
		return nil, errors.New("a valid website URL is required")
	}
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return nil, errors.New("website URL must use http or https")
	}
	return parsedURL, nil
}

func (c *Checker) httpClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}
	return http.DefaultClient
}

func (c *Checker) resolver() DNSResolver {
	if c.Resolver != nil {
		return c.Resolver
	}
	return net.DefaultResolver
}
