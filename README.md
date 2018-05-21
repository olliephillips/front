# Front

WORK IN PROGRESS

Outputs boilerplate web3.js compatible JavaScript by compiling from solidity or reading smart contract ABI. 

There's some [sample output here](https://github.com/olliephillips/front/blob/master/sample.js). Normally, generated output will be commented out, comments are omitted here for syntax highlighting.

## Install

It's a Go package so `go get` and `go install` it. When more complete I'll add binaries.

## Usage 

```
front --h
```

```
Usage of ./front:
  -abi string
        Path to the Ethereum contract ABI json to convert
  -address string
        Address of the contract/contracts
  -async
        Use asynchronous callbacks in with state changing functions
  -out string
        Output file for the generated web3.js javascript (default = stdout)
  -sol string
        Path to the Ethereum contract Solidity source to build and convert
  -solc string
        Solidity compiler to use if source builds are requested (default "solc")
```

If some of that looks familiar its because I wanted to give the application similar feel to Ethereum's abigen.


## Status

- Constructor, functions & events all done
- Testing to do