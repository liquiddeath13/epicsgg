package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func doBattle(u user, opponentID int, eventID int, stageID int) battleResultDescription {
	battle := battleDescription{}
	battle.BannedMapIds = []int{u.MapsForBan[0].ID, u.MapsForBan[1].ID}
	battle.CategoryID = 1
	battle.Circuit = circuitDescription{
		ID:      eventID,
		StageID: stageID,
	}
	battle.EnemyRosterID = opponentID
	battle.RosterID = u.RUSHRoster.ID
	content, _ := json.Marshal(battle)
	response := doRequest(host+"/ultimate-team/pve/games?", "POST", u.JWT, bytes.NewBuffer(content))
	var result battleResultDescription
	responseContent, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseContent, &result)
	return result
}
