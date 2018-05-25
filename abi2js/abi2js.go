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
func Convert(name string, abi string, async *bool, addr *string) (string, error) {
	var err error
	var js string

	var artifacts []artifact
	json.Unmarshal([]byte(abi), &artifacts)

	// need the constructor at this point, want it at the top
	// before the functions & events in the output
	for i := range artifacts {
		if artifacts[i].Type == "constructor" {
			val, err := parse(name, &artifacts[i], async, addr)
			if err != nil {
				fmt.Printf("Failed to parse ABI artifact: %v\n", err)
				os.Exit(-1)
			}
			//ok
			js += val
		}
	}

	for i := range artifacts {
		if artifacts[i].Type != "constructor" {
			val, err := parse(name, &artifacts[i], async, addr)
			if err != nil {
				fmt.Printf("Failed to parse ABI artifact: %v\n", err)
				os.Exit(-1)
			}
			// ok
			js += val
		}
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

// IncludeByteCode sets up a variable with the contract bytecode
// required to deploy a new contract via constructor
func IncludeByteCode(name string, code string) string {
	var js string
	js = fmt.Sprintf(byteCodeSyntax, name, name, code)
	return js
}

func parse(name string, art *artifact, async *bool, addr *string) (string, error) {
	// identify what this artifact is, and hand off
	var js string
	var err error

	switch art.Type {
	case "constructor":
		js, err = isConstructor(name, art, addr)
	case "function":
		js, err = isFunction(name, art, async)
	case "event":
		// events are always async
		js, err = isEvent(name, art)
	}
	return js, err
}

func camelCase(input string) string {
	// utility func for making camelCase strings
	return strings.ToLower(string(input[0])) + input[1:]
}
