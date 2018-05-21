package abi2js

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func isConstructor(name string, art *artifact, addr *string) (string, error) {
	var js string
	var err error

	js += initContract(name, art, addr)

	return js, err
}

func initContract(name string, art *artifact, addr *string) string {
	// InitContract includes JavaScript to initialise contracts based on ABI
	// If address is passed it will create code which initialises on ABI and address
	// If no address InitContract creates boilerplate to create from ABI and constructor
	var contractName, camelName, js string
	var address []byte
	var inputs, inputArgs string
	address = common.FromHex(*addr)
	contractName = name + "Contract"
	camelName = strings.ToLower(string(name[0])) + name[1:]

	if *addr != "" {
		// existing from address
		js = fmt.Sprintf(initContractFromAddress, contractName, contractName, name, camelName, contractName, camelName, contractName, address)
	} else {
		// new contract, code to create and use
		if len(art.Inputs) != 0 {
			inputs = "It accepts "
			for i := range art.Inputs {
				// comment
				inputs += art.Inputs[i].(map[string]interface{})["name"].(string) + " type "
				inputs += art.Inputs[i].(map[string]interface{})["type"].(string) + ", "
				// args
				inputArgs += art.Inputs[i].(map[string]interface{})["name"].(string) + ", "
			}
			inputs = strings.TrimRight(inputs, ", ") + "."
			inputs = "\n// " + inputs
		}
		js = fmt.Sprintf(initNewContract, contractName, contractName, camelName, camelName, contractName, inputs, camelName, camelName, name, camelName, contractName, inputArgs, camelName, camelName, camelName)
	}

	return js
}
