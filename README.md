# How to
```sh 
make install_dev_dep # to install mockgen
make install
make generate
make test 
make build 
./delio-take-home-test -stocks=AAPL,MSFT
```

# Assumptions & Explanations
- The requirement of calculating profit/loss comparing to the previous day's closing price is not necessary 
  - That information is automatically being pulled in from the `Quote` endpoint in field `d`
- The only calculations which I've had to do is 10x the stock values to get info on 10 stocks rather than one
  - As an example I have performed them concurrently but in this case it does not make sense to do this
  - It probably takes more time to spin up the go routines than it would take to 10x those values
  - However, I still done them concurrently just as an example

## Testing methodology
- I was having problems trying to mock the actual API library because I did not know how (if even possible) I could dependency inject `finnhub.ApiQuoteRequest` because it is a struct.
- Because of this I had to wrap and abstract the library away so that parts of the codebase which are dependent on it would be able to mock it in testing.
- As it comes to testing methodology I went with what I like to call verbose testing.
  - This is as opposed to abstracting tests away into struct arrays  to lessen the code duplication.
  - I agree that its nice to have code duplications, however, writing and fixing tests should be as easy for a developer as possible to encourage better code coverage.
  - Having the tests broken down into their individual function makes it far easier to read the test, amend them, fix them and write more tests
  - E.G. When a test fails, usually you get a line number where it fails, the struct abstraction can make it harder to find the failing test and the reason for it, causing a dev to waste more time of tests
