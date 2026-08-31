# Go Projects

This repository contains two separate Go projects for learning and practicing Go.

## Projects

### Booking-app

`Booking-app` is a command-line conference booking application. It collects attendee details, validates the input, tracks the number of tickets remaining, and displays the first names from the bookings.

This project is intended to be an executable program and should eventually be run from its directory with `go run .`. It currently needs its helper package layout corrected before it will build: `main.go` imports `Booking-app/helper`, but `helper.go` is currently in the project root.

### private-iphelper

`private-iphelper` is a website health-checking library. It resolves a website's DNS name, sends an HTTP or HTTPS request, measures response time, records redirects, and reports whether the website appears healthy.

This project is a library, so it does not run by itself with `go run`. From its directory, use `go test ./...` to build the package and run its tests. Use `go vet ./...` for static checks.

## Project layout

The two directories are independent Go modules. Run commands from the directory of the project you want to work on:

- `Booking-app`: command-line application.
- `private-iphelper`: reusable website-checking package.

Each project has its own `go.mod` file and dependency settings.
