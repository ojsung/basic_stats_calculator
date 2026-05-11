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
- [ ] Add binomial probability calculator
  - [x] Add binomial coefficient calculator
  - [ ] Add probability of success calculator
  - [ ] Create math_utils package to hold useful functions not implemented in `math/big`
    - [x] Factorial
    - [x] Power (a^n)
    - [x] Exp (e^x)
    - [ ] Ln
      - [x] Calculate using Taylor expansion for `ln(x+1)` and mantissa with base 2
      - [ ] Calculate edges of Taylor expansion's domain using Pade approximant
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
- [ ] Add CLI entrypoint to calculator

### V2
- [ ] Add p-value Calculator
- [ ] Add Cumulative Distribution Calculator
- [ ] Add HTTP server entrypoint to calculator
- [ ] Add basic html front-end for calculator
  - It seems like there's a common practice of hosting those straight out of related servers and tightly coupling servers to UIs. I may do this just for practice, since I feel okay about the process of hosting a web app separately

### V3
- [ ] Add a database to hold all the magic numbers
  - This is unnecessary. Holding them in the server is fine. But I want some practice working in Go with db's and this is the only thing I could think of to host
- [ ] Add a cache layer to hold... some numbers I guess
  - Same as above, unnecessary but I'd like some practice working with a caching layer. Probably Redis cause it's what I'm most familiar

### V4
- [ ] Add linear and exponential (power) regression calculator

## Current Tasks
### V0
- [ ] Switch math_utils.Ln to use Pade Approximant near those points instead of Taylor series expansion
   - It's currently using the Taylor expansion series to estimate Ln values between 0 and 2. However, at the edges of its domain (very near 0 and very near 2), the series converges _very_ slowly (for `ln(2)`, it will take around 2.5 million terms to reach 16 decimal places of accuracy)
   - [ ] Switch this to use a Pade Approximant
     - [x] Add matrix operations to allow programmatically solving linear equations
     - [ ] As needed, solve Pade Approximant terms to meet user's precision requirements
- [ ] Implement `calculator.calculateProbabilityOfKSuccesses`
  - Depends on Ln
- [ ] Implement `calculator.CalculateBinomialProbability`
  - Depends on `calculator.calculateProbabilityOfKSuccesses`
- [ ] Plan out CLI entrypoint
