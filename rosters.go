package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getCircuitRosterDetails(u user, rosterID int) rosterDescription {
	response := doRequest(fmt.Sprintf("%s/ultimate-team/pve/rosters?categoryId=1&ids=%d", host, rosterID), "GET", u.JWT, nil)
	content, _ := ioutil.ReadAll(response.Body)
	var rd rostersDescription
	json.Unmarshal(content, &rd)
	return rd.Data.Rosters[0]
}

func getUserRosters(userID int, jwt string) rostersDescription {
	response := doRequest(fmt.Sprintf("%s/ultimate-team/rosters?categoryId=1&userId=%d", host, userID), "GET", jwt, nil)
	content, _ := ioutil.ReadAll(response.Body)
	var rd rostersDescription
	json.Unmarshal(content, &rd)
	return rd
}
