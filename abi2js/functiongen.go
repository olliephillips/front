package abi2js

import (
	"fmt"
	"strings"
)

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
