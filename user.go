package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"time"
)

func retrieveAccountInfo(jwt string) accountInfoResponse {
	response := doRequest(host+"/user/info?", "GET", jwt, nil)
	content, _ := ioutil.ReadAll(response.Body)
	var me accountInfoResponse
	if debug {
		ioutil.WriteFile("me.json", content, 0644)
	}
	json.Unmarshal(content, &me)
	me.Data.JWT = jwt
	return me
}

func initUser(jwt string) user {
	accountInfo := retrieveAccountInfo(jwt)
	user := accountInfo.Data
	myRosters := getUserRosters(user.ID, user.JWT)
	allMaps := getMaps(user)
	maxRating := 0
	selectedRoster := rosterDescription{}
	for _, rosterDetail := range myRosters.Data.Rosters {
		if rosterDetail.Rating > maxRating {
			maxRating = rosterDetail.Rating
			selectedRoster = rosterDetail
		}
	}
	mapsForBan := []mapDescription{}
	m := selectedRoster.Stats.Maps
	sort.Slice(m, func(i, j int) bool {
		return m[i].Weight < m[j].Weight
	})
	mapForBan := allMaps.Data.Maps[m[1].MapID-1]
	mapForBan.RosterWeight = m[1].Weight
	mapsForBan = append(mapsForBan, mapForBan)
	mapForBan = allMaps.Data.Maps[m[0].MapID-1]
	mapForBan.RosterWeight = m[0].Weight
	mapsForBan = append(mapsForBan, mapForBan)
	user.RUSHRoster = selectedRoster
	user.MapsForBan = mapsForBan
	return user
}

func (u user) printInfo() {
	println("username:", u.Username)
	println("created:", u.Created.Format(time.RFC822))
	println("email:", u.Email)
	println("is email verified:", u.VerifiedEmail)
	println("balance (gold count):", u.Balance)
	println("is banned:", u.Banned)
}
