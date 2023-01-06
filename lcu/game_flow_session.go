package lcu

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gutil"
	"sort"
)

type SummonerListInGame struct {
	FriendTeam []Team `json:"friendTeam"`
	EnemyTeam  []Team `json:"enemyTeam"`
}

var PositionToIndex = map[string]int64{
	"BOTTOM":  4,
	"JUNGLE":  2,
	"TOP":     1,
	"MIDDLE":  3,
	"UTILITY": 5,
	"NONE":    0,
}

// GetEnemySummonerIdAndSummonerName 查询双方召唤师ID和昵称
func (c *Client) GetEnemySummonerIdAndSummonerName() SummonerListInGame {
	bs, _ := c.Do("GET", fmt.Sprintf("/lol-gameflow/v1/session"), nil)
	ret := toAny(bs, GameFlowSession{})
	// 英雄名称-id映射
	nameToId := make(map[string]int64, 10)
	for _, v := range ret.GameData.PlayerChampionSelections {
		nameToId[v.SummonerInternalName] = v.ChampionId
	}
	// 填充team.ChampId字段
	fillTeam(ret.GameData.TeamOne, nameToId, PositionToIndex)
	fillTeam(ret.GameData.TeamTwo, nameToId, PositionToIndex)
	// 友方/敌方
	var friendTeam, enemyTeam []Team
	if existAccount(CurrSummoner.SummonerID, ret.GameData.TeamOne) {
		friendTeam = ret.GameData.TeamOne
		enemyTeam = ret.GameData.TeamTwo
	} else {
		friendTeam = ret.GameData.TeamTwo
		enemyTeam = ret.GameData.TeamOne
	}
	gutil.Dump(friendTeam)
	return SummonerListInGame{FriendTeam: friendTeam, EnemyTeam: enemyTeam}
}

func fillTeam(list []Team, nameToId, positionToId map[string]int64) {
	for i := range list {
		list[i].ChampionId = nameToId[list[i].SummonerInternalName]
		list[i].PositionId = positionToId[list[i].SelectedPosition]
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].PositionId < list[j].PositionId
	})
}

func existAccount(id int64, list []Team) bool {
	for _, v := range list {
		if v.AccountId == id {
			return true
		}
	}
	return false
}

type GameFlowSession struct {
	GameClient GameClient `json:"gameClient"`
	GameData   GameData   `json:"gameData"`
	GameDodge  GameDodge  `json:"gameDodge"`
	Map        Map        `json:"map"`
	Phase      string     `json:"phase"`
}
type GameClient struct {
	ObserverServerIP   string `json:"observerServerIp"`
	ObserverServerPort int    `json:"observerServerPort"`
	Running            bool   `json:"running"`
	ServerIP           string `json:"serverIp"`
	ServerPort         int    `json:"serverPort"`
	Visible            bool   `json:"visible"`
}

type PlayerChampionSelections struct {
	AdditionalProp1      AdditionalProp1 `json:"additionalProp1"`
	SummonerInternalName string          `json:"summonerInternalName"`
	ChampionId           int64           `json:"championId"`
}
type GameTypeConfig struct {
	AdvancedLearningQuests bool   `json:"advancedLearningQuests"`
	AllowTrades            bool   `json:"allowTrades"`
	BanMode                string `json:"banMode"`
	BanTimerDuration       int    `json:"banTimerDuration"`
	BattleBoost            bool   `json:"battleBoost"`
	CrossTeamChampionPool  bool   `json:"crossTeamChampionPool"`
	DeathMatch             bool   `json:"deathMatch"`
	DoNotRemove            bool   `json:"doNotRemove"`
	DuplicatePick          bool   `json:"duplicatePick"`
	ExclusivePick          bool   `json:"exclusivePick"`
	ID                     int    `json:"id"`
	LearningQuests         bool   `json:"learningQuests"`
	MainPickTimerDuration  int    `json:"mainPickTimerDuration"`
	MaxAllowableBans       int    `json:"maxAllowableBans"`
	Name                   string `json:"name"`
	OnboardCoopBeginner    bool   `json:"onboardCoopBeginner"`
	PickMode               string `json:"pickMode"`
	PostPickTimerDuration  int    `json:"postPickTimerDuration"`
	Reroll                 bool   `json:"reroll"`
	TeamChampionPool       bool   `json:"teamChampionPool"`
}
type QueueRewards struct {
	IsChampionPointsEnabled bool  `json:"isChampionPointsEnabled"`
	IsIPEnabled             bool  `json:"isIpEnabled"`
	IsXpEnabled             bool  `json:"isXpEnabled"`
	PartySizeIPRewards      []int `json:"partySizeIpRewards"`
}
type Queue struct {
	AllowablePremadeSizes               []int          `json:"allowablePremadeSizes"`
	AreFreeChampionsAllowed             bool           `json:"areFreeChampionsAllowed"`
	AssetMutator                        string         `json:"assetMutator"`
	Category                            string         `json:"category"`
	ChampionsRequiredToPlay             int            `json:"championsRequiredToPlay"`
	Description                         string         `json:"description"`
	DetailedDescription                 string         `json:"detailedDescription"`
	GameMode                            string         `json:"gameMode"`
	GameTypeConfig                      GameTypeConfig `json:"gameTypeConfig"`
	ID                                  int            `json:"id"`
	IsRanked                            bool           `json:"isRanked"`
	IsTeamBuilderManaged                bool           `json:"isTeamBuilderManaged"`
	IsTeamOnly                          bool           `json:"isTeamOnly"`
	LastToggledOffTime                  int            `json:"lastToggledOffTime"`
	LastToggledOnTime                   int            `json:"lastToggledOnTime"`
	MapID                               int            `json:"mapId"`
	MaxLevel                            int            `json:"maxLevel"`
	MaxSummonerLevelForFirstWinOfTheDay int            `json:"maxSummonerLevelForFirstWinOfTheDay"`
	MaximumParticipantListSize          int            `json:"maximumParticipantListSize"`
	MinLevel                            int            `json:"minLevel"`
	MinimumParticipantListSize          int            `json:"minimumParticipantListSize"`
	Name                                string         `json:"name"`
	NumPlayersPerTeam                   int            `json:"numPlayersPerTeam"`
	QueueAvailability                   string         `json:"queueAvailability"`
	QueueRewards                        QueueRewards   `json:"queueRewards"`
	RemovalFromGameAllowed              bool           `json:"removalFromGameAllowed"`
	RemovalFromGameDelayMinutes         int            `json:"removalFromGameDelayMinutes"`
	ShortName                           string         `json:"shortName"`
	ShowPositionSelector                bool           `json:"showPositionSelector"`
	SpectatorEnabled                    bool           `json:"spectatorEnabled"`
	Type                                string         `json:"type"`
}
type Team struct {
	AdditionalProp1      AdditionalProp1 `json:"additionalProp1"`
	AccountId            int64           `json:"accountId"`
	SummonerName         string          `json:"summonerName"`
	SummonerInternalName string          `json:"summonerInternalName"`
	SelectedPosition     string          `json:"selectedPosition"`

	ChampionId int64 `json:"championId"`
	PositionId int64 `json:"positionId"`
}

type GameData struct {
	GameID                   int                        `json:"gameId"`
	GameName                 string                     `json:"gameName"`
	IsCustomGame             bool                       `json:"isCustomGame"`
	Password                 string                     `json:"password"`
	PlayerChampionSelections []PlayerChampionSelections `json:"playerChampionSelections"`
	Queue                    Queue                      `json:"queue"`
	SpectatorsAllowed        bool                       `json:"spectatorsAllowed"`
	TeamOne                  []Team                     `json:"teamOne"`
	TeamTwo                  []Team                     `json:"teamTwo"`
}
type GameDodge struct {
	DodgeIds []int  `json:"dodgeIds"`
	Phase    string `json:"phase"`
	State    string `json:"state"`
}
type Assets struct {
	AdditionalProp1 AdditionalProp1 `json:"additionalProp1"`
}
type CategorizedContentBundles struct {
	AdditionalProp1 AdditionalProp1 `json:"additionalProp1"`
}

type AdditionalProp1 struct {
	Spells []int `json:"spells"`
}
type AdditionalProp2 struct {
	Spells []int `json:"spells"`
}
type AdditionalProp3 struct {
	Spells []int `json:"spells"`
}
type PerPositionDisallowedSummonerSpells struct {
	AdditionalProp1 AdditionalProp1 `json:"additionalProp1"`
	AdditionalProp2 AdditionalProp2 `json:"additionalProp2"`
	AdditionalProp3 AdditionalProp3 `json:"additionalProp3"`
}
type PerPositionRequiredSummonerSpells struct {
	AdditionalProp1 AdditionalProp1 `json:"additionalProp1"`
	AdditionalProp2 AdditionalProp2 `json:"additionalProp2"`
	AdditionalProp3 AdditionalProp3 `json:"additionalProp3"`
}
type Properties struct {
	AdditionalProp1 AdditionalProp1 `json:"additionalProp1"`
}
type Map struct {
	Assets                              Assets                              `json:"assets"`
	CategorizedContentBundles           CategorizedContentBundles           `json:"categorizedContentBundles"`
	Description                         string                              `json:"description"`
	GameMode                            string                              `json:"gameMode"`
	GameModeName                        string                              `json:"gameModeName"`
	GameModeShortName                   string                              `json:"gameModeShortName"`
	GameMutator                         string                              `json:"gameMutator"`
	ID                                  int                                 `json:"id"`
	IsRGM                               bool                                `json:"isRGM"`
	MapStringID                         string                              `json:"mapStringId"`
	Name                                string                              `json:"name"`
	PerPositionDisallowedSummonerSpells PerPositionDisallowedSummonerSpells `json:"perPositionDisallowedSummonerSpells"`
	PerPositionRequiredSummonerSpells   PerPositionRequiredSummonerSpells   `json:"perPositionRequiredSummonerSpells"`
	PlatformID                          string                              `json:"platformId"`
	PlatformName                        string                              `json:"platformName"`
	Properties                          Properties                          `json:"properties"`
}
