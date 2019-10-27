# StockCollector

# Development Setup

## Setting up Environment
* Clone the project `git clone https://github.com/ryannel/StockCollector.git`
* [download golang](https://golang.org/dl/) 1.13 or greater and install with defaults
* Ensure that golang exists on your PATH by executing `go version` from the terminal
* [Generate an Alpha Vantage API key](https://www.alphavantage.co/support/#api-key) if you don't already have one
* Create an environment variable called `ALPHA_VANTAGE_API_KEY` and set your API key as the value

## Building the project
Execute the following commands in your terminal from the project root directory:

* Command line interface tool: `go build -o stockCollectorCli.exe ./1-inbound/cli/.`
