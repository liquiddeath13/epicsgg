package main

import "time"

type user struct {
	JWT                  string            `json:"-"`
	MapsForBan           []mapDescription  `json:"-"`
	RUSHRoster           rosterDescription `json:"-"`
	ID                   int               `json:"id"`
	FirstName            interface{}       `json:"firstName"`
	LastName             interface{}       `json:"lastName"`
	Username             string            `json:"username"`
	Email                string            `json:"email"`
	Phone                interface{}       `json:"phone"`
	Balance              int               `json:"balance"`
	Avatar               string            `json:"avatar"`
	EthAddress           interface{}       `json:"ethAddress"`
	VerifiedEmail        bool              `json:"verifiedEmail"`
	VerifiedPhone        bool              `json:"verifiedPhone"`
	VerifiedEthAddress   bool              `json:"verifiedEthAddress"`
	SteamID              interface{}       `json:"steamId"`
	SteamTradeCode       interface{}       `json:"steamTradeCode"`
	Country              string            `json:"country"`
	Banned               bool              `json:"banned"`
	Group                int               `json:"group"`
	PackPublic           bool              `json:"packPublic"`
	ChatPushNotification bool              `json:"chatPushNotification"`
	Created              time.Time         `json:"created"`
	Onboarding           bool              `json:"onboarding"`
	Steam                interface{}       `json:"steam"`
	Twitch               interface{}       `json:"twitch"`
	Circuit              struct {
		Stage string `json:"stage"`
		Badge string `json:"badge"`
	} `json:"circuit"`
}

type accountInfoResponse struct {
	Success bool `json:"success"`
	Data    user `json:"data"`
}

type battleDescription struct {
	RosterID      int                `json:"rosterId"`
	EnemyRosterID int                `json:"enemyRosterId"`
	BannedMapIds  []int              `json:"bannedMapIds"`
	Circuit       circuitDescription `json:"circuit"`
	CategoryID    int                `json:"categoryId"`
}

type circuitDescription struct {
	ID      int `json:"id"`
	StageID int `json:"stageId"`
}

type battleResultDescription struct {
	Success bool `json:"success"`
	Data    struct {
		Game struct {
			ID     int    `json:"id"`
			MapID  int    `json:"mapId"`
			Replay string `json:"replay"`
			User1  struct {
				ID       int    `json:"id"`
				Username string `json:"username"`
				Avatar   string `json:"avatar"`
				Winner   bool   `json:"winner"`
				Side     string `json:"side"`
				Roster   struct {
					Name   string `json:"name"`
					Rating int    `json:"rating"`
					Level  int    `json:"level"`
					Stars  int    `json:"stars"`
					Cards  []struct {
						RoleID    int  `json:"roleId"`
						IsFlex    bool `json:"isFlex"`
						CardLevel int  `json:"cardLevel"`
						Card      struct {
							ID     int    `json:"id"`
							Status string `json:"status"`
							UUID   string `json:"uuid"`
							Images struct {
								Size402 string `json:"size402"`
								Svg     struct {
									Size200 string `json:"size200"`
									Size400 string `json:"size400"`
									Size600 string `json:"size600"`
									Full    string `json:"full"`
								} `json:"svg"`
							} `json:"images"`
							EthStatus         string      `json:"ethStatus"`
							Type              string      `json:"type"`
							MintNumber        int         `json:"mintNumber"`
							MintBatch         string      `json:"mintBatch"`
							Rating            float64     `json:"rating"`
							Level             int         `json:"level"`
							LevelBonus        int         `json:"levelBonus"`
							SignatureImage    interface{} `json:"signatureImage"`
							IsNewCardTemplate bool        `json:"isNewCardTemplate"`
							IsGhost           bool        `json:"isGhost"`
							IsMarketList      bool        `json:"isMarketList"`
							IsTradeList       bool        `json:"isTradeList"`
							CardTemplate      struct {
								ID          int    `json:"id"`
								UUID        string `json:"uuid"`
								Title       string `json:"title"`
								CardType    string `json:"cardType"`
								CategoryID  int    `json:"categoryId"`
								TreatmentID int    `json:"treatmentId"`
								Properties  struct {
									Salary       int    `json:"salary"`
									Season       string `json:"season"`
									GameID       int    `json:"game_id"`
									TeamID       int    `json:"team_id"`
									DateEnd      string `json:"date_end"`
									PlayerID     int    `json:"player_id"`
									DateStart    string `json:"date_start"`
									PlayerRating int    `json:"player_rating"`
								} `json:"properties"`
								LimitedEdition bool   `json:"limitedEdition"`
								InCirculation  int    `json:"inCirculation"`
								Rarity         string `json:"rarity"`
								Images         struct {
									Size102 time.Time `json:"size102"`
									Size201 time.Time `json:"size201"`
									Size402 time.Time `json:"size402"`
									Svg     struct {
										Size200 time.Time `json:"size200"`
										Size400 time.Time `json:"size400"`
										Size600 time.Time `json:"size600"`
										Full    time.Time `json:"full"`
									} `json:"svg"`
								} `json:"images"`
								Game struct {
									ID        int    `json:"id"`
									Name      string `json:"name"`
									DateStart string `json:"dateStart"`
									DateEnd   string `json:"dateEnd"`
									Image     string `json:"image"`
									Images    []struct {
										ID          string      `json:"id"`
										ParentType  string      `json:"parentType"`
										ParentID    int         `json:"parentId"`
										Name        string      `json:"name"`
										Position    string      `json:"position"`
										URL         string      `json:"url"`
										CardSide    string      `json:"cardSide"`
										TreatmentID interface{} `json:"treatmentId"`
										Properties  interface{} `json:"properties"`
									} `json:"images"`
								} `json:"game"`
								Team struct {
									ID      int    `json:"id"`
									Country string `json:"country"`
									Dob     string `json:"dob"`
									Name    string `json:"name"`
									GameID  int    `json:"gameId"`
									Active  bool   `json:"active"`
									Image   string `json:"image"`
									Images  []struct {
										ID          string      `json:"id"`
										ParentType  string      `json:"parentType"`
										ParentID    int         `json:"parentId"`
										Name        string      `json:"name"`
										Position    string      `json:"position"`
										URL         string      `json:"url"`
										CardSide    string      `json:"cardSide"`
										TreatmentID int         `json:"treatmentId"`
										Properties  interface{} `json:"properties"`
									} `json:"images"`
									Bio       string `json:"bio"`
									DateStart string `json:"dateStart"`
									DateEnd   string `json:"dateEnd"`
									ShortName string `json:"shortName"`
									Manager   string `json:"manager"`
									PlayerIds []int  `json:"playerIds"`
								} `json:"team"`
								Player struct {
									ID      int    `json:"id"`
									Dob     string `json:"dob"`
									Age     int    `json:"age"`
									Country string `json:"country"`
									Name    string `json:"name"`
									GameID  int    `json:"gameId"`
									TeamID  int    `json:"teamId"`
									Handle  string `json:"handle"`
									Active  bool   `json:"active"`
									Image   string `json:"image"`
									Images  []struct {
										ID          string      `json:"id"`
										ParentType  string      `json:"parentType"`
										ParentID    int         `json:"parentId"`
										Name        string      `json:"name"`
										Position    string      `json:"position"`
										URL         string      `json:"url"`
										CardSide    string      `json:"cardSide"`
										TreatmentID interface{} `json:"treatmentId"`
										Properties  interface{} `json:"properties"`
									} `json:"images"`
									Videos []struct {
										ID          int    `json:"id"`
										ParentType  string `json:"parentType"`
										ParentID    int    `json:"parentId"`
										Name        string `json:"name"`
										Position    string `json:"position"`
										URL         string `json:"url"`
										CardSide    string `json:"cardSide"`
										TreatmentID int    `json:"treatmentId"`
									} `json:"videos"`
									Bio          string `json:"bio"`
									Position     string `json:"position"`
									DateStart    string `json:"dateStart"`
									DateEnd      string `json:"dateEnd"`
									Quote        string `json:"quote"`
									HomeTown     string `json:"homeTown"`
									FrameType    string `json:"frameType"`
									LastDate     string `json:"lastDate"`
									PlayerRoleID int    `json:"playerRoleId"`
								} `json:"player"`
								PlayerStats []struct {
									Name  string  `json:"name"`
									Value float64 `json:"value"`
								} `json:"playerStats"`
								PlayerStatsV2 struct {
									Rating struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"rating"`
									Accuracy struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"accuracy"`
									Impact struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"impact"`
									Assists struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"assists"`
									Entry struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"entry"`
									Utility struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"utility"`
									Experience struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"experience"`
								} `json:"playerStatsV2"`
								PlayerRole struct {
									Name         string `json:"name"`
									GameSideIcon string `json:"gameSideIcon"`
								} `json:"playerRole"`
								Level struct {
									Costs struct {
										Epicoins    interface{} `json:"epicoins"`
										Silvercoins struct {
											Num2 int `json:"2"`
											Num3 int `json:"3"`
											Num4 int `json:"4"`
											Num5 int `json:"5"`
										} `json:"silvercoins"`
									} `json:"costs"`
									Bonuses struct {
										Num1 int     `json:"1"`
										Num2 float64 `json:"2"`
										Num3 float64 `json:"3"`
										Num4 float64 `json:"4"`
										Num5 float64 `json:"5"`
									} `json:"bonuses"`
								} `json:"level"`
								Treatment struct {
									ID          int         `json:"id"`
									Name        string      `json:"name"`
									CategoryID  int         `json:"categoryId"`
									Designation string      `json:"designation"`
									Tier        string      `json:"tier"`
									Active      bool        `json:"active"`
									Variation   string      `json:"variation"`
									GameSide    interface{} `json:"gameSide"`
									AccentColor interface{} `json:"accentColor"`
									ArtistName  interface{} `json:"artistName"`
									Season      string      `json:"season"`
									BuyPrice    interface{} `json:"buyPrice"`
									Images      []struct {
										ID          string      `json:"id"`
										ParentType  string      `json:"parentType"`
										ParentID    int         `json:"parentId"`
										Name        string      `json:"name"`
										Position    interface{} `json:"position"`
										URL         string      `json:"url"`
										CardSide    string      `json:"cardSide"`
										TreatmentID interface{} `json:"treatmentId"`
										Properties  interface{} `json:"properties"`
									} `json:"images"`
									Videos         []interface{} `json:"videos"`
									SuggestedPrice interface{}   `json:"suggestedPrice"`
								} `json:"treatment"`
							} `json:"cardTemplate"`
							EthTransactions []interface{} `json:"ethTransactions"`
						} `json:"card"`
					} `json:"cards"`
					MapBonus         float64 `json:"mapBonus"`
					CardLevelBonus   int     `json:"cardLevelBonus"`
					TeamChemistry    float64 `json:"teamChemistry"`
					CountryChemistry float64 `json:"countryChemistry"`
					RoleChemistry    float64 `json:"roleChemistry"`
				} `json:"roster"`
				BannedMapIds []int   `json:"bannedMapIds"`
				WinChance    float64 `json:"winChance"`
				Points       int     `json:"points"`
			} `json:"user1"`
			User2 struct {
				ID       interface{} `json:"id"`
				Username interface{} `json:"username"`
				Avatar   interface{} `json:"avatar"`
				Winner   bool        `json:"winner"`
				Side     string      `json:"side"`
				Roster   struct {
					Name   string `json:"name"`
					Rating int    `json:"rating"`
					Level  int    `json:"level"`
					Stars  int    `json:"stars"`
					Cards  []struct {
						RoleID    int  `json:"roleId"`
						IsFlex    bool `json:"isFlex"`
						CardLevel int  `json:"cardLevel"`
						Card      struct {
							ID     int    `json:"id"`
							Status string `json:"status"`
							UUID   string `json:"uuid"`
							Images struct {
								Size402 string `json:"size402"`
								Svg     struct {
									Size200 string `json:"size200"`
									Size400 string `json:"size400"`
									Size600 string `json:"size600"`
									Full    string `json:"full"`
								} `json:"svg"`
							} `json:"images"`
							EthStatus         string      `json:"ethStatus"`
							Type              string      `json:"type"`
							MintNumber        int         `json:"mintNumber"`
							MintBatch         string      `json:"mintBatch"`
							Rating            int         `json:"rating"`
							Level             int         `json:"level"`
							LevelBonus        int         `json:"levelBonus"`
							SignatureImage    interface{} `json:"signatureImage"`
							IsNewCardTemplate bool        `json:"isNewCardTemplate"`
							IsGhost           bool        `json:"isGhost"`
							IsMarketList      bool        `json:"isMarketList"`
							IsTradeList       bool        `json:"isTradeList"`
							CardTemplate      struct {
								ID          int    `json:"id"`
								UUID        string `json:"uuid"`
								Title       string `json:"title"`
								CardType    string `json:"cardType"`
								CategoryID  int    `json:"categoryId"`
								TreatmentID int    `json:"treatmentId"`
								Properties  struct {
									Salary       int    `json:"salary"`
									Season       string `json:"season"`
									GameID       int    `json:"game_id"`
									TeamID       int    `json:"team_id"`
									DateEnd      string `json:"date_end"`
									PlayerID     int    `json:"player_id"`
									DateStart    string `json:"date_start"`
									PlayerRating int    `json:"player_rating"`
								} `json:"properties"`
								LimitedEdition bool   `json:"limitedEdition"`
								InCirculation  int    `json:"inCirculation"`
								Rarity         string `json:"rarity"`
								Images         struct {
									Size102 time.Time `json:"size102"`
									Size201 time.Time `json:"size201"`
									Size402 time.Time `json:"size402"`
									Svg     struct {
										Size200 time.Time `json:"size200"`
										Size400 time.Time `json:"size400"`
										Size600 time.Time `json:"size600"`
										Full    time.Time `json:"full"`
									} `json:"svg"`
								} `json:"images"`
								Game struct {
									ID        int    `json:"id"`
									Name      string `json:"name"`
									DateStart string `json:"dateStart"`
									DateEnd   string `json:"dateEnd"`
									Image     string `json:"image"`
									Images    []struct {
										ID          string      `json:"id"`
										ParentType  string      `json:"parentType"`
										ParentID    int         `json:"parentId"`
										Name        string      `json:"name"`
										Position    string      `json:"position"`
										URL         string      `json:"url"`
										CardSide    string      `json:"cardSide"`
										TreatmentID interface{} `json:"treatmentId"`
										Properties  interface{} `json:"properties"`
									} `json:"images"`
								} `json:"game"`
								Team struct {
									ID      int    `json:"id"`
									Country string `json:"country"`
									Dob     string `json:"dob"`
									Name    string `json:"name"`
									GameID  int    `json:"gameId"`
									Active  bool   `json:"active"`
									Image   string `json:"image"`
									Images  []struct {
										ID          string      `json:"id"`
										ParentType  string      `json:"parentType"`
										ParentID    int         `json:"parentId"`
										Name        string      `json:"name"`
										Position    string      `json:"position"`
										URL         string      `json:"url"`
										CardSide    string      `json:"cardSide"`
										TreatmentID interface{} `json:"treatmentId"`
										Properties  interface{} `json:"properties"`
									} `json:"images"`
									Bio       string `json:"bio"`
									DateStart string `json:"dateStart"`
									DateEnd   string `json:"dateEnd"`
									ShortName string `json:"shortName"`
									Manager   string `json:"manager"`
									PlayerIds []int  `json:"playerIds"`
								} `json:"team"`
								Player struct {
									ID      int    `json:"id"`
									Dob     string `json:"dob"`
									Age     int    `json:"age"`
									Country string `json:"country"`
									Name    string `json:"name"`
									GameID  int    `json:"gameId"`
									TeamID  int    `json:"teamId"`
									Handle  string `json:"handle"`
									Active  bool   `json:"active"`
									Image   string `json:"image"`
									Images  []struct {
										ID          string      `json:"id"`
										ParentType  string      `json:"parentType"`
										ParentID    int         `json:"parentId"`
										Name        string      `json:"name"`
										Position    string      `json:"position"`
										URL         string      `json:"url"`
										CardSide    string      `json:"cardSide"`
										TreatmentID interface{} `json:"treatmentId"`
										Properties  interface{} `json:"properties"`
									} `json:"images"`
									Videos       []interface{} `json:"videos"`
									Bio          string        `json:"bio"`
									Position     string        `json:"position"`
									DateStart    string        `json:"dateStart"`
									DateEnd      string        `json:"dateEnd"`
									Quote        string        `json:"quote"`
									HomeTown     string        `json:"homeTown"`
									FrameType    string        `json:"frameType"`
									LastDate     string        `json:"lastDate"`
									PlayerRoleID int           `json:"playerRoleId"`
								} `json:"player"`
								PlayerStats []struct {
									Name  string  `json:"name"`
									Value float64 `json:"value"`
								} `json:"playerStats"`
								PlayerStatsV2 struct {
									Rating struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"rating"`
									Accuracy struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"accuracy"`
									Impact struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"impact"`
									Assists struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"assists"`
									Entry struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"entry"`
									Utility struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"utility"`
									Experience struct {
										Name  string `json:"name"`
										Short string `json:"short"`
										Score int    `json:"score"`
									} `json:"experience"`
								} `json:"playerStatsV2"`
								PlayerRole struct {
									Name         string `json:"name"`
									GameSideIcon string `json:"gameSideIcon"`
								} `json:"playerRole"`
								Level struct {
									Costs struct {
										Epicoins    interface{} `json:"epicoins"`
										Silvercoins struct {
											Num2 int `json:"2"`
											Num3 int `json:"3"`
											Num4 int `json:"4"`
											Num5 int `json:"5"`
										} `json:"silvercoins"`
									} `json:"costs"`
									Bonuses struct {
										Num1 int     `json:"1"`
										Num2 float64 `json:"2"`
										Num3 float64 `json:"3"`
										Num4 float64 `json:"4"`
										Num5 float64 `json:"5"`
									} `json:"bonuses"`
								} `json:"level"`
								Treatment struct {
									ID          int         `json:"id"`
									Name        string      `json:"name"`
									CategoryID  int         `json:"categoryId"`
									Designation string      `json:"designation"`
									Tier        string      `json:"tier"`
									Active      bool        `json:"active"`
									Variation   string      `json:"variation"`
									GameSide    interface{} `json:"gameSide"`
									AccentColor interface{} `json:"accentColor"`
									ArtistName  interface{} `json:"artistName"`
									Season      string      `json:"season"`
									BuyPrice    interface{} `json:"buyPrice"`
									Images      []struct {
										ID          string      `json:"id"`
										ParentType  string      `json:"parentType"`
										ParentID    int         `json:"parentId"`
										Name        string      `json:"name"`
										Position    interface{} `json:"position"`
										URL         string      `json:"url"`
										CardSide    string      `json:"cardSide"`
										TreatmentID interface{} `json:"treatmentId"`
										Properties  interface{} `json:"properties"`
									} `json:"images"`
									Videos         []interface{} `json:"videos"`
									SuggestedPrice interface{}   `json:"suggestedPrice"`
								} `json:"treatment"`
							} `json:"cardTemplate"`
							EthTransactions []interface{} `json:"ethTransactions"`
						} `json:"card"`
					} `json:"cards"`
					MapBonus         float64 `json:"mapBonus"`
					CardLevelBonus   int     `json:"cardLevelBonus"`
					TeamChemistry    float64 `json:"teamChemistry"`
					CountryChemistry float64 `json:"countryChemistry"`
					RoleChemistry    int     `json:"roleChemistry"`
				} `json:"roster"`
				BannedMapIds []int   `json:"bannedMapIds"`
				WinChance    float64 `json:"winChance"`
				Points       int     `json:"points"`
			} `json:"user2"`
		} `json:"game"`
	} `json:"data"`
}

type rostersDescription struct {
	Success bool `json:"success"`
	Data    struct {
		Rosters []rosterDescription `json:"rosters"`
	} `json:"data"`
}

type rosterDescription struct {
	ID              int               `json:"id"`
	CategoryID      int               `json:"categoryId"`
	Rating          int               `json:"rating"`
	Level           int               `json:"level"`
	Salary          int               `json:"salary"`
	SalaryOverLimit bool              `json:"salaryOverLimit"`
	Stars           int               `json:"stars"`
	Name            string            `json:"name"`
	Cards           []cardDescription `json:"cards"`
	Stats           rosterStats       `json:"stats"`
}

type rosterStats struct {
	Position         float64         `json:"position"`
	CountryChemistry float64         `json:"countryChemistry"`
	TeamChemistry    float64         `json:"teamChemistry"`
	CardLevel        int             `json:"cardLevel"`
	Maps             []rosterMapStat `json:"maps"`
}

type rosterMapStat struct {
	MapID  int     `json:"mapId"`
	Weight float64 `json:"weight"`
}

type cardDescription struct {
	RoleID    int  `json:"roleId"`
	IsFlex    bool `json:"isFlex"`
	CardLevel int  `json:"cardLevel"`
	Card      struct {
		ID                int         `json:"id"`
		Status            string      `json:"status"`
		UUID              string      `json:"uuid"`
		EthStatus         string      `json:"ethStatus"`
		Type              string      `json:"type"`
		MintNumber        int         `json:"mintNumber"`
		MintBatch         string      `json:"mintBatch"`
		Rating            float64     `json:"rating"`
		Level             int         `json:"level"`
		LevelBonus        int         `json:"levelBonus"`
		SignatureImage    interface{} `json:"signatureImage"`
		IsNewCardTemplate bool        `json:"isNewCardTemplate"`
		IsGhost           bool        `json:"isGhost"`
		IsMarketList      bool        `json:"isMarketList"`
		IsTradeList       bool        `json:"isTradeList"`
		CardTemplate      struct {
			ID          int    `json:"id"`
			UUID        string `json:"uuid"`
			Title       string `json:"title"`
			CardType    string `json:"cardType"`
			CategoryID  int    `json:"categoryId"`
			TreatmentID int    `json:"treatmentId"`
			Properties  struct {
				Salary       int    `json:"salary"`
				Season       string `json:"season"`
				GameID       int    `json:"game_id"`
				TeamID       int    `json:"team_id"`
				DateEnd      string `json:"date_end"`
				PlayerID     int    `json:"player_id"`
				DateStart    string `json:"date_start"`
				PlayerRating int    `json:"player_rating"`
			} `json:"properties"`
			LimitedEdition bool   `json:"limitedEdition"`
			InCirculation  int    `json:"inCirculation"`
			Rarity         string `json:"rarity"`
			Game           struct {
				ID        int    `json:"id"`
				Name      string `json:"name"`
				DateStart string `json:"dateStart"`
				DateEnd   string `json:"dateEnd"`
				Image     string `json:"image"`
			} `json:"game"`
			Team struct {
				ID        int    `json:"id"`
				Country   string `json:"country"`
				Dob       string `json:"dob"`
				Name      string `json:"name"`
				GameID    int    `json:"gameId"`
				Active    bool   `json:"active"`
				Image     string `json:"image"`
				Bio       string `json:"bio"`
				DateStart string `json:"dateStart"`
				DateEnd   string `json:"dateEnd"`
				ShortName string `json:"shortName"`
				Manager   string `json:"manager"`
				PlayerIds []int  `json:"playerIds"`
			} `json:"team"`
			Player struct {
				ID           int    `json:"id"`
				Dob          string `json:"dob"`
				Age          int    `json:"age"`
				Country      string `json:"country"`
				Name         string `json:"name"`
				GameID       int    `json:"gameId"`
				TeamID       int    `json:"teamId"`
				Handle       string `json:"handle"`
				Active       bool   `json:"active"`
				Image        string `json:"image"`
				Bio          string `json:"bio"`
				Position     string `json:"position"`
				DateStart    string `json:"dateStart"`
				DateEnd      string `json:"dateEnd"`
				Quote        string `json:"quote"`
				HomeTown     string `json:"homeTown"`
				FrameType    string `json:"frameType"`
				LastDate     string `json:"lastDate"`
				PlayerRoleID int    `json:"playerRoleId"`
			} `json:"player"`
			PlayerStats []struct {
				Name  string  `json:"name"`
				Value float64 `json:"value"`
			} `json:"playerStats"`
			PlayerStatsV2 struct {
				Rating struct {
					Name  string `json:"name"`
					Short string `json:"short"`
					Score int    `json:"score"`
				} `json:"rating"`
				Accuracy struct {
					Name  string `json:"name"`
					Short string `json:"short"`
					Score int    `json:"score"`
				} `json:"accuracy"`
				Impact struct {
					Name  string `json:"name"`
					Short string `json:"short"`
					Score int    `json:"score"`
				} `json:"impact"`
				Assists struct {
					Name  string `json:"name"`
					Short string `json:"short"`
					Score int    `json:"score"`
				} `json:"assists"`
				Entry struct {
					Name  string `json:"name"`
					Short string `json:"short"`
					Score int    `json:"score"`
				} `json:"entry"`
				Utility struct {
					Name  string `json:"name"`
					Short string `json:"short"`
					Score int    `json:"score"`
				} `json:"utility"`
				Experience struct {
					Name  string `json:"name"`
					Short string `json:"short"`
					Score int    `json:"score"`
				} `json:"experience"`
			} `json:"playerStatsV2"`
			PlayerRole struct {
				Name         string `json:"name"`
				GameSideIcon string `json:"gameSideIcon"`
			} `json:"playerRole"`
			Level struct {
				Costs struct {
					Epicoins    interface{} `json:"epicoins"`
					Silvercoins struct {
						Num2 int `json:"2"`
						Num3 int `json:"3"`
						Num4 int `json:"4"`
						Num5 int `json:"5"`
					} `json:"silvercoins"`
				} `json:"costs"`
				Bonuses struct {
					Num1 int     `json:"1"`
					Num2 float64 `json:"2"`
					Num3 float64 `json:"3"`
					Num4 float64 `json:"4"`
					Num5 float64 `json:"5"`
				} `json:"bonuses"`
			} `json:"level"`
			Treatment struct {
				ID             int         `json:"id"`
				Name           string      `json:"name"`
				CategoryID     int         `json:"categoryId"`
				Designation    string      `json:"designation"`
				Tier           string      `json:"tier"`
				Active         bool        `json:"active"`
				Variation      string      `json:"variation"`
				GameSide       interface{} `json:"gameSide"`
				AccentColor    interface{} `json:"accentColor"`
				ArtistName     interface{} `json:"artistName"`
				Season         string      `json:"season"`
				BuyPrice       interface{} `json:"buyPrice"`
				SuggestedPrice interface{} `json:"suggestedPrice"`
			} `json:"treatment"`
		} `json:"cardTemplate"`
		EthTransactions []interface{} `json:"ethTransactions"`
	} `json:"card"`
}

type mapList struct {
	Success bool `json:"success"`
	Data    struct {
		Maps []mapDescription `json:"maps"`
	} `json:"data"`
}

type mapDescription struct {
	RosterWeight float64 `json:"-"`
	ID           int     `json:"id"`
	CategoryID   int     `json:"categoryId"`
	Name         string  `json:"name"`
	Pve          bool    `json:"pve"`
	Pvp          bool    `json:"pvp"`
}

type circuitDetails struct {
	Success bool `json:"success"`
	Data    struct {
		Circuit struct {
			ID              int         `json:"id"`
			ParentCircuitID interface{} `json:"parentCircuitId"`
			PrestigeOrder   interface{} `json:"prestigeOrder"`
			Name            string      `json:"name"`
			Description     string      `json:"description"`
			IsCore          bool        `json:"isCore"`
			Start           time.Time   `json:"start"`
			End             time.Time   `json:"end"`
			StagesCompleted int         `json:"stagesCompleted"`
			TotalStages     int         `json:"totalStages"`
			Properties      struct {
				AccentColor string `json:"accent_color"`
			} `json:"properties"`
			Rules struct {
				RosterLevel int `json:"rosterLevel"`
				CardLevel   struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"cardLevel"`
			} `json:"rules"`
			Stages []struct {
				ID        int         `json:"id"`
				Name      string      `json:"name"`
				Objective string      `json:"objective"`
				Completed interface{} `json:"completed"`
				IsClaimed bool        `json:"isClaimed"`
				Rewards   struct {
					Coins         int           `json:"coins"`
					Silvercoins   int           `json:"silvercoins"`
					Craftingcoins int           `json:"craftingcoins"`
					CardTemplates []interface{} `json:"card_templates"`
					PackTemplates []struct {
						ID       int `json:"id"`
						Quantity int `json:"quantity"`
					} `json:"pack_templates"`
				} `json:"rewards"`
				Rosters []struct {
					Wins          int `json:"wins"`
					UtPveRosterID int `json:"ut_pve_roster_id"`
				} `json:"rosters"`
				RosterProgress []struct {
					UtPveRosterID int `json:"ut_pve_roster_id"`
					Wins          int `json:"wins"`
					Losses        int `json:"losses"`
					MatchWins     int `json:"match_wins"`
					MatchLosses   int `json:"match_losses"`
				} `json:"rosterProgress"`
			} `json:"stages"`
			AllStages []struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Order int    `json:"order"`
				Badge string `json:"badge"`
			} `json:"allStages"`
		} `json:"circuit"`
	} `json:"data"`
}

type eventList struct {
	Success bool `json:"success"`
	Data    struct {
		Circuits []struct {
			ID              int         `json:"id"`
			PrestigeOrder   interface{} `json:"prestigeOrder"`
			ParentCircuitID interface{} `json:"parentCircuitId"`
			Name            string      `json:"name"`
			Description     string      `json:"description"`
			IsCore          bool        `json:"isCore"`
			Start           time.Time   `json:"start"`
			End             time.Time   `json:"end"`
			Rules           struct {
				RosterLevel int `json:"rosterLevel"`
				CardLevel   struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"cardLevel"`
			} `json:"rules"`
			StagesCompleted int `json:"stagesCompleted"`
			TotalStages     int `json:"totalStages"`
			Properties      struct {
				AccentColor string `json:"accent_color"`
			} `json:"properties"`
			Stages []struct {
				ID        int         `json:"id"`
				Name      string      `json:"name"`
				Objective string      `json:"objective"`
				Completed interface{} `json:"completed"`
				IsClaimed bool        `json:"isClaimed"`
				Rewards   struct {
					Coins         int           `json:"coins"`
					Silvercoins   int           `json:"silvercoins"`
					Craftingcoins int           `json:"craftingcoins"`
					CardTemplates []interface{} `json:"card_templates"`
					PackTemplates []struct {
						ID       int `json:"id"`
						Quantity int `json:"quantity"`
					} `json:"pack_templates"`
				} `json:"rewards"`
			} `json:"stages"`
			AllStages []struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Order int    `json:"order"`
				Badge string `json:"badge"`
			} `json:"allStages"`
		} `json:"circuits"`
	} `json:"data"`
}
