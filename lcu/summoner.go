package lcu

import (
	"encoding/json"
	"fmt"
	"log"
)

type Summoner struct {
	AccountID                   int64        `json:"accountId"`
	DisplayName                 string       `json:"displayName"`
	InternalName                string       `json:"internalName"`
	NameChangeFlag              bool         `json:"nameChangeFlag"`
	PercentCompleteForNextLevel int          `json:"percentCompleteForNextLevel"`
	Privacy                     string       `json:"privacy"`
	ProfileIconID               int          `json:"profileIconId"`
	Puuid                       string       `json:"puuid"`
	RerollPoints                RerollPoints `json:"rerollPoints"`
	SummonerID                  int64        `json:"summonerId"`
	SummonerLevel               int          `json:"summonerLevel"`
	Unnamed                     bool         `json:"unnamed"`
	XpSinceLastLevel            int          `json:"xpSinceLastLevel"`
	XpUntilNextLevel            int          `json:"xpUntilNextLevel"`
}

type RerollPoints struct {
	CurrentPoints    int `json:"currentPoints"`
	MaxRolls         int `json:"maxRolls"`
	NumberOfRolls    int `json:"numberOfRolls"`
	PointsCostToRoll int `json:"pointsCostToRoll"`
	PointsToReroll   int `json:"pointsToReroll"`
}

// GetCurrentSummoner 获取当前召唤师信息
func (c *Client) GetCurrentSummoner() Summoner {
	bs, _ := c.Do("GET", fmt.Sprintf("/lol-summoner/v1/current-summoner"), nil)
	ret := toAny(bs, Summoner{})
	if ret.SummonerID != 0 {
		CurrSummoner = &ret
	}
	return ret
}

var CurrSummoner *Summoner

// GetSummoner 根据id获取召唤师信息
func (c *Client) GetSummoner(id string) Summoner {
	bs, _ := c.Do("GET", fmt.Sprintf("/lol-summoner/v1/summoners/%s", id), nil)
	ret := toAny(bs, Summoner{})
	return ret
}

func (c *Client) GetSummonerMastery(id string) Summoner {
	bs, _ := c.Do("GET", fmt.Sprintf("/lol-collections/v1/inventories/%s/champion-mastery", id), nil)
	ret := toAny(bs, Summoner{})
	return ret
}

func toAny[T any](bs []byte, c T) T {
	err := json.Unmarshal(bs, &c)
	if err != nil {
		log.Println("toAny failed", err, "bs:", string(bs))
		return c
	}
	return c
}
