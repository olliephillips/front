package abi2js

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

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

func parse(name string, art *artifact, async *bool) (string, error) {
	var js string
	var err error

	switch art.Type {
	case "function":
		js, err = isFunction(name, art, async)
	case "event":
		//js, err = isEvent(art)
	case "constructor":
		//js, err = isConstructor(art)
	}
	return js, err
}

func isFunction(name string, art *artifact, async *bool) (string, error) {
	var js string
	var err error

	// make comment
	js += makeFunctionComment(art)

	// make body
	js += makeFunctionBody(name, art, async)

	return js, err
}

func makeFunctionComment(art *artifact) string {
	var comment string
	var typ, inputs, outputs, trnsNote string
	if len(art.Inputs) != 0 {
		inputs = "It accepts "
		for i := range art.Inputs {
			inputs += art.Inputs[i].(map[string]interface{})["name"].(string) + " type "
			inputs += art.Inputs[i].(map[string]interface{})["type"].(string) + ", "
		}
		inputs = strings.TrimRight(inputs, ", ") + "."
	}

	if len(art.Outputs) != 0 {
		outputs = "It returns "
		for i := range art.Outputs {
			outputs += art.Outputs[i].Name
			if art.Outputs[i].Name == "" {
				outputs += "type "
			} else {
				outputs += " type "
			}
			outputs += art.Outputs[i].Type + ", "
		}
		outputs = strings.TrimRight(outputs, ", ") + "."
	}

	// operations dependent on function type
	switch art.StateMutability {
	case "nonpayable":
		typ = ""
		trnsNote = "\n// Transaction object parameter 'gas' is wei denominated"
	case "payable":
		typ = " " + art.StateMutability
		trnsNote = "\n// Transaction object parameters, 'value' and 'gas' are wei denominated"
	case "view":
		typ = " " + art.StateMutability
	}

	comment = fmt.Sprintf(functionComment, art.Name, typ, inputs, outputs, trnsNote)

	return comment
}

func makeFunctionBody(name string, art *artifact, async *bool) string {
	var res, instance, method, inputArgs, trnsObj, callback, js string
	res = camelCase(art.Name) + "Res"
	instance = camelCase(name)
	method = art.Name

	if len(art.Inputs) != 0 {
		for i := range art.Inputs {
			name := art.Inputs[i].(map[string]interface{})["name"].(string)
			inputArgs += name + ", "
		}
	}

	// operations dependent on function type
	switch art.StateMutability {
	case "nonpayable":
		trnsObj = "{gas:0}"

	case "payable":
		trnsObj = "{value: 0, gas: 0}"

	case "view":
		trnsObj = ""
		inputArgs = strings.TrimRight(inputArgs, ", ")
	}

	if art.StateMutability == "view" {
		js = fmt.Sprintf(viewFunctionBody, res, instance, method, inputArgs)
	} else {
		// we need to pass transaction value and gas here
		if *async {
			// callbacks - to do
			callback = callbackAsync
			js = fmt.Sprintf(functionBodyAsync, instance, method, inputArgs, trnsObj, callback)
		} else {
			js = fmt.Sprintf(functionBody, instance, method, inputArgs, trnsObj)
		}
	}
	return js
}

func isEvent() {
	// not implemented yet

}

func isConstructor() {
	// not implemented yet
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

// InitContract includes JavaScrip to initialise contracts based on ABI
func InitContract(name string) string {
	var contractName, camelName, js string
	contractName = name + "Contract"
	camelName = strings.ToLower(string(name[0])) + name[1:]
	js = fmt.Sprintf(initContract, contractName, contractName, name, camelName, contractName, camelName, contractName)
	return js
}

func camelCase(input string) string {
	return strings.ToLower(string(input[0])) + input[1:]
}
