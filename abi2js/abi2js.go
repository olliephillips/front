package abi2js

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// struct for unmarshaling the ABI
type artifact struct {
	Constant bool          `json:"constant,omitempty"`
	Inputs   []interface{} `json:"inputs"`
	Name     string        `json:"name,omitempty"`
	Outputs  []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"outputs,omitempty"`
	Payable         bool   `json:"payable,omitempty"`
	StateMutability string `json:"stateMutability,omitempty"`
	Type            string `json:"type"`
	Anonymous       bool   `json:"anonymous,omitempty"`
}

// Convert takes contract ABI in JSON format and
// converts it to web3.js compatible javascript code
func Convert(name string, abi string, async *bool) (string, error) {
	var err error
	var js string

	var artifacts []artifact
	json.Unmarshal([]byte(abi), &artifacts)

	for i := range artifacts {
		val, err := parse(name, &artifacts[i], async)
		if err != nil {
			fmt.Printf("Failed to parse ABI artifact: %v\n", err)
			os.Exit(-1)
		}
		// ok
		js += val
	}

	return js, err
}

// InitWeb3 includes JavaScript to initialise web provider,
// or use current provider
func InitWeb3() string {
	var js string
	js = web3Init
	return js
}

// IncludeABI sets up the ABI as a variable in the JavaScript
func IncludeABI(name string, abi string) string {
	var js string
	js = fmt.Sprintf(abiSyntax, name, name, abi)
	return js
}

// InitContract includes JavaScript to initialise contracts based on ABI
func InitContract(name string) string {
	var contractName, camelName, js string
	contractName = name + "Contract"
	camelName = strings.ToLower(string(name[0])) + name[1:]
	js = fmt.Sprintf(initContract, contractName, contractName, name, camelName, contractName, camelName, contractName)
	return js
}

func parse(name string, art *artifact, async *bool) (string, error) {
	// identify what this artifact is, and hand off
	var js string
	var err error

	switch art.Type {
	case "function":
		js, err = isFunction(name, art, async)
	case "event":
		// events are alway async
		js, err = isEvent(name, art)
	case "constructor":
		//js, err = isConstructor(art)
	}
	return js, err
}

func camelCase(input string) string {
	// utility func for making camelCase strings
	return strings.ToLower(string(input[0])) + input[1:]
}

/*
func isConstructor() {
	// not implemented yet
}
*/
