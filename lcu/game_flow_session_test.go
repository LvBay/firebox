package lcu

import (
	"testing"
)

func TestClient_GetEnemySummonerIdAndSummonerName(t *testing.T) {
	LocalTestClient.GetCurrentSummoner()
	LocalTestClient.GetEnemySummonerIdAndSummonerName()
}
