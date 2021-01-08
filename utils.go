package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func getMaps(u user) mapList {
	response := doRequest(fmt.Sprintf("%s/ultimate-team/maps?categoryId=1", host), "GET", u.JWT, nil)
	content, _ := ioutil.ReadAll(response.Body)
	var ml mapList
	json.Unmarshal(content, &ml)
	return ml
}
