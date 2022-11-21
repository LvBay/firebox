package lcu

import "testing"

func TestGetCurrentSummoner(t *testing.T) {
	LocalTestClient.GetCurrentSummoner()
}

func TestGetSummoner(t *testing.T) {
	LocalTestClient.GetSummoner("2951059970")
}

func TestGetSummonerMastery(t *testing.T) {
	LocalTestClient.GetSummonerMastery("2951059970")
}

var LocalTestClient *Client

func init() {
	var err error
	LocalTestClient, err = NewClient()
	if err != nil {
		panic("lol未启动")
	}
}
