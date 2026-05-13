# basic_stats_calculator

A calculator with some basic statistics functions. Very basic. Written in Go so I can get some practice in with the language
Trying to keep the structure as Clean as possible. Trying to keep it as TDD and tested as possible

## Getting Started

There's not currently any usable entrypoints to the application. Once there's at least one (CLI entrypoint is planned for V1), I will update this section. For now, if you want, you can run the tests

- Clone this repository
- Install `go` for your system
- Run `go test ./...`

## Roadmap

A tentative, short-term roadmap. I don't intend for this to be a long project, since I'm just doing it for fun. Roadmaps for release versions will become more specific as I start researching and planning them

### V1 (minimum viable product)

- [x] Add binomial probability calculator
  - [x] Add binomial coefficient calculator
  - [x] Add probability of success calculator
  - [x] Create math_utils package to hold useful functions not implemented in `math/big`
    - [x] Factorial
    - [x] Power (a^n)
    - [x] Exp (e^x)
    - [x] Ln
      - [x] Calculate using Taylor expansion for `ln(x+1)` and mantissa with base 2
      - [x] Calculate edges of Taylor expansion's domain using Pade approximant
        - [x] Add matrix operations to allow programmatically solving linear equations
          - [x] Determinant
          - [x] Cofactor
          - [x] Cofactor Matrix
          - [x] Transpose
          - [x] Inverse
          - [x] Trace
          - [x] Add
          - [x] Subtract
          - [x] Scalar Multiply
          - [x] Matrix Multiply
          - [x] Remove columns and rows
    - [x] Euler's Number
  - [x] Create big_utils package to hold useful functions for working with values in `math/big`
    - [x] String to Float
    - [x] Float to String
    - [x] Compare Floats to accuracy point
- [x] Add HTTP/HTML entrypoint to calculator

### V2

- [ ] Add p-value Calculator
- [ ] Add Cumulative Distribution Calculator


### V3

- [ ] Add a database to hold all the magic numbers
  - This is unnecessary. Holding them in the server is fine. But I want some practice working in Go with db's and this is the only thing I could think of to host
- [ ] Add a cache layer to hold... some numbers I guess
  - Same as above, unnecessary but I'd like some practice working with a caching layer. Probably Redis cause it's what I'm most familiar

### V4

- [ ] Add linear and exponential (power) regression calculator

## Current Tasks

None. V1 is complete. I will pick this back up for V2-V4 if I ever feel like it