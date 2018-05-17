package abi2js

var web3Init = `
// Initialise new Web3 provider or use current provider
if (typeof web3 !== 'undefined') {
	web3 = new Web3(web3.currentProvider);
} else {
	// set the provider you want from Web3.providers
	web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
}
`
var functionComment = `

// %s is a%s function. %s %s%s 
`
var abiSyntax = `
// %s contract ABI
var %sABI = %s
`
var initContract = `
// %s is a contract object
var %s = web3.eth.contract(%sABI);

// %s is an instance of the %s contract object
var %s = %s.at('CONTRACTADDRESSHERE');`

var viewFunctionBody = `var %s = %s.%s(%s);`

var functionBody = `%s.%s(%s%s);`

var functionBodyAsync = `%s.%s(%s%s);`
