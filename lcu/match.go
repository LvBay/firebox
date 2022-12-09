package lcu

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *Client) GetMatchInfo(gameID string) (MatchInfo, error) {
	bs, _ := c.Do("GET", fmt.Sprintf("/lol-match-history/v1/games/%s", gameID), nil)
	var ret MatchInfo
	err := json.Unmarshal(bs, &ret)
	fillStats(&ret)
	return ret, err
}

func fillStats(info *MatchInfo) {
	m := map[int]*Stats{}
	for _, v := range info.Participants {
		m[v.ParticipantID] = &v.Stats
	}
	for i, v := range info.Participants {
		info.ParticipantIdentities[i].Stats = m[v.ParticipantID]
	}
}

func (c *Client) GetMatchListBySummonerId(summonerId string) MatchList {
	bs, _ := c.Do("GET", fmt.Sprintf("/lol-match-history/v3/matchlist/account/%s", summonerId), nil)
	ret := toAny(bs, MatchList{})
	for i := range ret.Games.Games {
		fillStats(&ret.Games.Games[i])
	}
	return ret
}

type MatchInfo struct {
	GameCreation          int64                   `json:"gameCreation"`
	GameCreationDate      time.Time               `json:"gameCreationDate"`
	GameDuration          int                     `json:"gameDuration"`
	GameID                int64                   `json:"gameId"`
	GameMode              string                  `json:"gameMode"`
	GameType              string                  `json:"gameType"`
	GameVersion           string                  `json:"gameVersion"`
	MapID                 int                     `json:"mapId"`
	ParticipantIdentities []ParticipantIdentities `json:"participantIdentities"`
	Participants          []Participants          `json:"participants"`
	PlatformID            string                  `json:"platformId"`
	QueueID               int                     `json:"queueId"`
	SeasonID              int                     `json:"seasonId"`
	Teams                 []Teams                 `json:"teams"`
}
type Player struct {
	AccountID         int64  `json:"accountId"`
	CurrentAccountID  int64  `json:"currentAccountId"`
	CurrentPlatformID string `json:"currentPlatformId"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	PlatformID        string `json:"platformId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerID        int64  `json:"summonerId"`
	SummonerName      string `json:"summonerName"`
}
type ParticipantIdentities struct {
	ParticipantID int    `json:"participantId"`
	Player        Player `json:"player"`
	Stats         *Stats `json:"stats"`
}
type Stats struct {
	Assists                         int  `json:"assists"`
	CausedEarlySurrender            bool `json:"causedEarlySurrender"`
	ChampLevel                      int  `json:"champLevel"`
	CombatPlayerScore               int  `json:"combatPlayerScore"`
	DamageDealtToObjectives         int  `json:"damageDealtToObjectives"`
	DamageDealtToTurrets            int  `json:"damageDealtToTurrets"`
	DamageSelfMitigated             int  `json:"damageSelfMitigated"`
	Deaths                          int  `json:"deaths"`
	DoubleKills                     int  `json:"doubleKills"`
	EarlySurrenderAccomplice        bool `json:"earlySurrenderAccomplice"`
	FirstBloodAssist                bool `json:"firstBloodAssist"`
	FirstBloodKill                  bool `json:"firstBloodKill"`
	FirstInhibitorAssist            bool `json:"firstInhibitorAssist"`
	FirstInhibitorKill              bool `json:"firstInhibitorKill"`
	FirstTowerAssist                bool `json:"firstTowerAssist"`
	FirstTowerKill                  bool `json:"firstTowerKill"`
	GameEndedInEarlySurrender       bool `json:"gameEndedInEarlySurrender"`
	GameEndedInSurrender            bool `json:"gameEndedInSurrender"`
	GoldEarned                      int  `json:"goldEarned"`
	GoldSpent                       int  `json:"goldSpent"`
	InhibitorKills                  int  `json:"inhibitorKills"`
	Item0                           int  `json:"item0"`
	Item1                           int  `json:"item1"`
	Item2                           int  `json:"item2"`
	Item3                           int  `json:"item3"`
	Item4                           int  `json:"item4"`
	Item5                           int  `json:"item5"`
	Item6                           int  `json:"item6"`
	KillingSprees                   int  `json:"killingSprees"`
	Kills                           int  `json:"kills"`
	LargestCriticalStrike           int  `json:"largestCriticalStrike"`
	LargestKillingSpree             int  `json:"largestKillingSpree"`
	LargestMultiKill                int  `json:"largestMultiKill"`
	LongestTimeSpentLiving          int  `json:"longestTimeSpentLiving"`
	MagicDamageDealt                int  `json:"magicDamageDealt"`
	MagicDamageDealtToChampions     int  `json:"magicDamageDealtToChampions"`
	MagicalDamageTaken              int  `json:"magicalDamageTaken"`
	NeutralMinionsKilled            int  `json:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int  `json:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  int  `json:"neutralMinionsKilledTeamJungle"`
	ObjectivePlayerScore            int  `json:"objectivePlayerScore"`
	ParticipantID                   int  `json:"participantId"`
	PentaKills                      int  `json:"pentaKills"`
	Perk0                           int  `json:"perk0"`
	Perk0Var1                       int  `json:"perk0Var1"`
	Perk0Var2                       int  `json:"perk0Var2"`
	Perk0Var3                       int  `json:"perk0Var3"`
	Perk1                           int  `json:"perk1"`
	Perk1Var1                       int  `json:"perk1Var1"`
	Perk1Var2                       int  `json:"perk1Var2"`
	Perk1Var3                       int  `json:"perk1Var3"`
	Perk2                           int  `json:"perk2"`
	Perk2Var1                       int  `json:"perk2Var1"`
	Perk2Var2                       int  `json:"perk2Var2"`
	Perk2Var3                       int  `json:"perk2Var3"`
	Perk3                           int  `json:"perk3"`
	Perk3Var1                       int  `json:"perk3Var1"`
	Perk3Var2                       int  `json:"perk3Var2"`
	Perk3Var3                       int  `json:"perk3Var3"`
	Perk4                           int  `json:"perk4"`
	Perk4Var1                       int  `json:"perk4Var1"`
	Perk4Var2                       int  `json:"perk4Var2"`
	Perk4Var3                       int  `json:"perk4Var3"`
	Perk5                           int  `json:"perk5"`
	Perk5Var1                       int  `json:"perk5Var1"`
	Perk5Var2                       int  `json:"perk5Var2"`
	Perk5Var3                       int  `json:"perk5Var3"`
	PerkPrimaryStyle                int  `json:"perkPrimaryStyle"`
	PerkSubStyle                    int  `json:"perkSubStyle"`
	PhysicalDamageDealt             int  `json:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  int  `json:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int  `json:"physicalDamageTaken"`
	PlayerScore0                    int  `json:"playerScore0"`
	PlayerScore1                    int  `json:"playerScore1"`
	PlayerScore2                    int  `json:"playerScore2"`
	PlayerScore3                    int  `json:"playerScore3"`
	PlayerScore4                    int  `json:"playerScore4"`
	PlayerScore5                    int  `json:"playerScore5"`
	PlayerScore6                    int  `json:"playerScore6"`
	PlayerScore7                    int  `json:"playerScore7"`
	PlayerScore8                    int  `json:"playerScore8"`
	PlayerScore9                    int  `json:"playerScore9"`
	QuadraKills                     int  `json:"quadraKills"`
	SightWardsBoughtInGame          int  `json:"sightWardsBoughtInGame"`
	TeamEarlySurrendered            bool `json:"teamEarlySurrendered"`
	TimeCCingOthers                 int  `json:"timeCCingOthers"`
	TotalDamageDealt                int  `json:"totalDamageDealt"`
	TotalDamageDealtToChampions     int  `json:"totalDamageDealtToChampions"`
	TotalDamageTaken                int  `json:"totalDamageTaken"`
	TotalHeal                       int  `json:"totalHeal"`
	TotalMinionsKilled              int  `json:"totalMinionsKilled"`
	TotalPlayerScore                int  `json:"totalPlayerScore"`
	TotalScoreRank                  int  `json:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int  `json:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int  `json:"totalUnitsHealed"`
	TripleKills                     int  `json:"tripleKills"`
	TrueDamageDealt                 int  `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      int  `json:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int  `json:"trueDamageTaken"`
	TurretKills                     int  `json:"turretKills"`
	UnrealKills                     int  `json:"unrealKills"`
	VisionScore                     int  `json:"visionScore"`
	VisionWardsBoughtInGame         int  `json:"visionWardsBoughtInGame"`
	WardsKilled                     int  `json:"wardsKilled"`
	WardsPlaced                     int  `json:"wardsPlaced"`
	Win                             bool `json:"win"`
}
type CreepsPerMinDeltas struct {
	Zero10    float64 `json:"0-10"`
	One020    float64 `json:"10-20"`
	Two030    float64 `json:"20-30"`
	Three0End float64 `json:"30-end"`
}
type CsDiffPerMinDeltas struct {
	Zero10    float64 `json:"0-10"`
	One020    float64 `json:"10-20"`
	Two030    float64 `json:"20-30"`
	Three0End float64 `json:"30-end"`
}
type DamageTakenDiffPerMinDeltas struct {
	Zero10    float64 `json:"0-10"`
	One020    float64 `json:"10-20"`
	Two030    float64 `json:"20-30"`
	Three0End float64 `json:"30-end"`
}
type DamageTakenPerMinDeltas struct {
	Zero10    float64 `json:"0-10"`
	One020    float64 `json:"10-20"`
	Two030    float64 `json:"20-30"`
	Three0End float64 `json:"30-end"`
}
type GoldPerMinDeltas struct {
	Zero10    float64 `json:"0-10"`
	One020    float64 `json:"10-20"`
	Two030    float64 `json:"20-30"`
	Three0End float64 `json:"30-end"`
}
type XpDiffPerMinDeltas struct {
	Zero10    float64 `json:"0-10"`
	One020    float64 `json:"10-20"`
	Two030    float64 `json:"20-30"`
	Three0End float64 `json:"30-end"`
}
type XpPerMinDeltas struct {
	Zero10    float64 `json:"0-10"`
	One020    float64 `json:"10-20"`
	Two030    float64 `json:"20-30"`
	Three0End float64 `json:"30-end"`
}
type Timeline struct {
	CreepsPerMinDeltas          CreepsPerMinDeltas          `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas          CsDiffPerMinDeltas          `json:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas DamageTakenDiffPerMinDeltas `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     DamageTakenPerMinDeltas     `json:"damageTakenPerMinDeltas"`
	GoldPerMinDeltas            GoldPerMinDeltas            `json:"goldPerMinDeltas"`
	Lane                        string                      `json:"lane"`
	ParticipantID               int                         `json:"participantId"`
	Role                        string                      `json:"role"`
	XpDiffPerMinDeltas          XpDiffPerMinDeltas          `json:"xpDiffPerMinDeltas"`
	XpPerMinDeltas              XpPerMinDeltas              `json:"xpPerMinDeltas"`
}
type Participants struct {
	ChampionID                int      `json:"championId"`
	HighestAchievedSeasonTier string   `json:"highestAchievedSeasonTier"`
	ParticipantID             int      `json:"participantId"`
	Spell1ID                  int      `json:"spell1Id"`
	Spell2ID                  int      `json:"spell2Id"`
	Stats                     Stats    `json:"stats"`
	TeamID                    int      `json:"teamId"`
	Timeline                  Timeline `json:"timeline"`
}
type Bans struct {
	ChampionID int `json:"championId"`
	PickTurn   int `json:"pickTurn"`
}
type Teams struct {
	Bans                 []Bans `json:"bans"`
	BaronKills           int    `json:"baronKills"`
	DominionVictoryScore int    `json:"dominionVictoryScore"`
	DragonKills          int    `json:"dragonKills"`
	FirstBaron           bool   `json:"firstBaron"`
	FirstBlood           bool   `json:"firstBlood"`
	FirstDargon          bool   `json:"firstDargon"`
	FirstInhibitor       bool   `json:"firstInhibitor"`
	FirstTower           bool   `json:"firstTower"`
	InhibitorKills       int    `json:"inhibitorKills"`
	RiftHeraldKills      int    `json:"riftHeraldKills"`
	TeamID               int    `json:"teamId"`
	TowerKills           int    `json:"towerKills"`
	VilemawKills         int    `json:"vilemawKills"`
	Win                  string `json:"win"`
}

// matchList
type MatchList struct {
	AccountID  int64  `json:"accountId"`
	Games      Games  `json:"games"`
	PlatformID string `json:"platformId"`
}

//	type GamesGames struct {
//		GameCreation          int64                   `json:"gameCreation"`
//		GameCreationDate      time.Time               `json:"gameCreationDate"`
//		GameDuration          int                     `json:"gameDuration"`
//		GameID                int64                   `json:"gameId"`
//		GameMode              string                  `json:"gameMode"`
//		GameType              string                  `json:"gameType"`
//		GameVersion           string                  `json:"gameVersion"`
//		MapID                 int                     `json:"mapId"`
//		ParticipantIdentities []ParticipantIdentities `json:"participantIdentities"`
//		Participants          []Participants          `json:"participants"`
//		PlatformID            string                  `json:"platformId"`
//		QueueID               int                     `json:"queueId"`
//		SeasonID              int                     `json:"seasonId"`
//		Teams                 []interface{}           `json:"teams"`
//	}

type Games struct {
	GameBeginDate  string      `json:"gameBeginDate"`
	GameCount      int         `json:"gameCount"`
	GameEndDate    string      `json:"gameEndDate"`
	GameIndexBegin int         `json:"gameIndexBegin"`
	GameIndexEnd   int         `json:"gameIndexEnd"`
	Games          []MatchInfo `json:"games"`
}
