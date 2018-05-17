/**
 ** Code generated - manual changes may be lost if regenerated
 ** Generated by Front, solc or abi to web3 javascript converter
 ** http://github.com/olliephillips/front
**/

// Initialise new Web3 provider or use current provider
if (typeof web3 !== 'undefined') {
        web3 = new Web3(web3.currentProvider);
} else {
        // set the provider you want from Web3.providers
        web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
}

// SavingsAccount contract ABI
var SavingsAccountABI = [{"constant":true,"inputs":[],"name":"GetOwner","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_initials","type":"string"}],"name":"ReceiveByPerson","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":false,"inputs":[{"name":"_initials","type":"string"},{"name":"_account","type":"address"}],"name":"ChangeAccount","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_initials","type":"string"}],"name":"GetAccountBalance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"GetBankBalance","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"KillContract","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"Receive","outputs":[],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"_initials","type":"string"}],"name":"GetAccount","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"ShareFallback","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_newOwner","type":"address"}],"name":"ChangeOwner","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[],"payable":true,"stateMutability":"payable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"address"},{"indexed":false,"name":"","type":"uint256"}],"name":"ContractDeposit","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"string"},{"indexed":false,"name":"","type":"address"},{"indexed":false,"name":"","type":"uint256"}],"name":"DirectDeposit","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"","type":"string"},{"indexed":false,"name":"","type":"uint256"}],"name":"Distribution","type":"event"}]

// SavingsAccountContract is a contract object
var SavingsAccountContract = web3.eth.contract(SavingsAccountABI);

// savingsAccount is an instance of the SavingsAccountContract contract object
var savingsAccount = SavingsAccountContract.at('CONTRACTADDRESSHERE');

// GetOwner is a view function.  It returns type address.
var getOwnerRes = savingsAccount.GetOwner();

// ReceiveByPerson is a payable function. It accepts _initials type string.
// Transaction object parameters, 'value' and 'gas' are wei denominated
savingsAccount.ReceiveByPerson(_initials, {value: 0, gas: 0});

// ChangeAccount is a function. It accepts _initials type string, _account type address.
// Transaction object parameter 'gas' is wei denominated
savingsAccount.ChangeAccount(_initials, _account, {gas:0});

// GetAccountBalance is a view function. It accepts _initials type string. It returns type uint256.
var getAccountBalanceRes = savingsAccount.GetAccountBalance(_initials);

// GetBankBalance is a view function.  It returns type uint256.
var getBankBalanceRes = savingsAccount.GetBankBalance();

// KillContract is a function.
// Transaction object parameter 'gas' is wei denominated
savingsAccount.KillContract({gas:0});

// Receive is a payable function.
// Transaction object parameters, 'value' and 'gas' are wei denominated
savingsAccount.Receive({value: 0, gas: 0});

// GetAccount is a view function. It accepts _initials type string. It returns type address.
var getAccountRes = savingsAccount.GetAccount(_initials);

// ShareFallback is a function.
// Transaction object parameter 'gas' is wei denominated
savingsAccount.ShareFallback({gas:0});

// ChangeOwner is a function. It accepts _newOwner type address.
// Transaction object parameter 'gas' is wei denominated
savingsAccount.ChangeOwner(_newOwner, {gas:0});