package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getCircuitDetails(u user, circuitID int) circuitDetails {
	response := doRequest(fmt.Sprintf("%s/ultimate-team/circuits/35", host), "GET", u.JWT, nil)
	content, _ := ioutil.ReadAll(response.Body)
	var cd circuitDetails
	json.Unmarshal(content, &cd)
	return cd
}

func getCircuits(u user) eventList {
	response := doRequest(fmt.Sprintf("%s/ultimate-team/circuits?categoryId=1", host), "GET", u.JWT, nil)
	content, _ := ioutil.ReadAll(response.Body)
	var el eventList
	json.Unmarshal(content, &el)
	return el
}
