package abi2js

import (
	"fmt"
	"strings"
)

func isEvent(name string, art *artifact) (string, error) {
	var js string
	var err error

	// make comment
	js += makeEventComment(art)

	// make body
	js += makeEventBody(name, art)

	return js, err
}

func makeEventComment(art *artifact) string {
	var comment string
	var data string //, trnsNote string
	if len(art.Inputs) != 0 {
		data = "It sends "
		for i := range art.Inputs {
			data += art.Inputs[i].(map[string]interface{})["name"].(string) + "type "
			data += art.Inputs[i].(map[string]interface{})["type"].(string) + ", "
		}
		data = strings.TrimRight(data, ", ") + " on occurence"
	}

	comment = fmt.Sprintf(eventComment, art.Name, data)
	return comment
}

func makeEventBody(name string, art *artifact) string {
	var evtName, conName, logName, js string
	evtName = camelCase(art.Name) + "Event"
	logName = camelCase(art.Name) + "EventLog"
	conName = camelCase(name)
	js = fmt.Sprintf(eventBody, art.Name, evtName, conName, art.Name, evtName, logName, evtName, evtName, evtName)
	return js
}
