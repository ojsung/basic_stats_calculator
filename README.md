# basic_stats_calculator

A calculator with some basic statistics functions. Very basic. Written in Go so I can get some practice in with the language.
Trying to keep the structure as clean as possible. Trying to keep it as TDD and tested as possible.

## Getting Started

### Running the tests

- Clone this repository
- Install `go` for your system
- Run `go test ./...`

### Running the web server

- Run `go run ./cmd/stats_server/` (or `make run`)
- Opens on `http://localhost:8080` by default
- Set the `PORT` environment variable to use a different port

| Route | What it does |
|---|---|
| `/` | Binomial probability — P(X = k) |
| `/cdf` | Cumulative distribution — P(X ≤ k), with per-term breakdown |
| `/pvalue` | Binomial p-value — left-tail, right-tail, or two-tail |

## Roadmap

### Released

**V1** — Binomial probability calculator with HTTP/HTML web server

**V2** — Cumulative distribution and p-value calculators (left/right/two-tail)

### V3 Break out packages

- [ ] Break out big_utils to own package
- [ ] Break out matrix + operand into own package

### V4

- [ ] Add linear and exponential (power) regression calculator

## Current Tasks

None. V2 is complete. I will pick this back up for V3 and V4 if I ever feel like it.
