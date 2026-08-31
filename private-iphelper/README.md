# private-iphelper

This is a Go library for checking whether a website is reachable and healthy. It resolves the website's hostname, sends an HTTP or HTTPS request, measures the response time, records redirects, and reports the HTTP status code.

A website is considered healthy when DNS resolution succeeds, the request succeeds, the response status is below 400, and an HTTPS connection has valid TLS information.

## How to run it

This project is a library and does not have its own executable command. Open a terminal in the `private-iphelper` directory and use the following commands:

- `go mod tidy` updates the module dependencies.
- `go test ./...` builds the package and runs the automated tests.
- `go vet ./...` runs additional Go checks.

Applications that check real websites need internet access and access to a DNS resolver. The automated tests use a local test server and do not depend on external websites.

## Files

### checker.go

This is the main package file. It contains the website checker, result type, DNS resolution, HTTP and HTTPS requests, redirect tracking, response-time measurement, and health-status logic.

### checker_test.go

This file contains automated tests. It tests a healthy website, a website returning an internal-server-error response, and invalid URL input. The tests use local test data and a fake DNS resolver so they are predictable and do not depend on external websites.

## Why are there two Go files?

The files have different responsibilities:

- `checker.go` contains the code developers use.
- `checker_test.go` contains code that verifies the package works correctly.

Go recognizes files ending in `_test.go` as test files and only includes them when running commands such as `go test`. The test code is not included in programs that import this package.

