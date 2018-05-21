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
var %sABI = %s;
`

var byteCodeSyntax = `
// %s contract bytecode
var %sByteCode = '%s';
`

var initContractFromAddress = `
// %s is a contract object
var %s = web3.eth.contract(%sABI);

// %s is an instance of the %s
// contract object, created using existing contract address
var %s = %s.at('0x%x');`

var initNewContract = `
// %s is a contract object
var %s = web3.eth.contract(%sABI);

// %s is an instance of the %s
// contract object, created as a new contract via Constructor.%s
var %sDeployAddress = web3.eth.accounts[0];
var %sGas = web3.eth.estimateGas({data: %sByteCode});
var %s = %s.new(%s{data: %sByteCode, from: %sDeployAddress, gas: %sGas});`

var viewFunctionBody = `var %s = %s.%s(%s);`

var functionBody = `%s.%s(%s%s);`

var functionBodyAsync = `%s.%s(%s%s%s);`

var callbackAsync = `, (err, res) => {
  if(!err)
    // console.log(JSON.stringify(res));
  else
   console.log(err);	
}`

var eventComment = `

// %s is an event. %s`

var eventBody = `
// init and watch for %s event
var %s = %s.%s({}, { fromBlock: 0, toBlock: 'latest' }, (err, evt) => {
  if (!err)
    console.log(evt);
  else 
    console.log(err);
});

// get entire log for %s
var %s = %s.get((err, log) => {
  if(!err)
    console.log(log);
});

// stop watching %s event
%s.stopWatching();`
