package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	me := initUser(token)
	me.printInfo()
	f, err := os.OpenFile("epicsgg.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println(fmt.Sprintf("start at %s", time.Now().Format(time.RFC3339)))
	log.Println(fmt.Sprintf("will ban: %s (%f), %s (%f)",
		me.MapsForBan[0].Name,
		me.MapsForBan[0].RosterWeight,
		me.MapsForBan[1].Name,
		me.MapsForBan[1].RosterWeight))
	events := getCircuits(me).Data.Circuits
	for _, event := range events {
		cd := getCircuitDetails(me, event.ID).Data.Circuit
		log.Println(fmt.Sprintf("%s in progress..", cd.Name))
		for _, stage := range cd.Stages {
			if stage.Completed == nil {
				for _, opponent := range stage.Rosters {
					opponentDetails := getCircuitRosterDetails(me, opponent.UtPveRosterID)
					log.Println(fmt.Sprintf("vs %s", opponentDetails.Name))
					if me.RUSHRoster.Rating < opponentDetails.Rating {
						panic("stopped: opponent rating more than our!")
					}
					winsCounter := 0
					for winsCounter < opponent.Wins {
						battleResult := doBattle(me, opponent.UtPveRosterID, 35, stage.ID)
						if battleResult.Data.Game.User1.Winner {
							points := battleResult.Data.Game.User1.Points
							if points > 0 {
								log.Println(fmt.Sprintf("win! +%d points", points))
								log.Println(battleResult.Data.Game.Replay)
								winsCounter++
							} else {
								//FIXME: IT DOES ONLY ONE BATTLE PER TEAM (CHECK BATTLE USER, MAYBE SECOND?)
								//log.Println("seems like we already destroyed that team. going to the next...")
								//break
							}
						}
						time.Sleep(5 * time.Second)
					}
				}
			}
		}
	}
}
