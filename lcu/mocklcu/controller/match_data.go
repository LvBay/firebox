package controller

import (
	"github.com/LvBay/firebox/lcu"
	"github.com/gogf/gf/v2/util/grand"
)

var yiVsGe = lcu.MatchInfo{
	GameDuration: 1800,
	GameID:       10001,
	ParticipantIdentities: []lcu.ParticipantIdentities{
		{
			ParticipantID: 1,
			Player: lcu.Player{
				SummonerName: "浪翻云",
			},
		},
		{
			ParticipantID: 2,
			Player: lcu.Player{
				SummonerName: "凌战天",
			},
		},
		{
			ParticipantID: 3,
			Player: lcu.Player{
				SummonerName: "干罗",
			},
		},
		{
			ParticipantID: 4,
			Player: lcu.Player{
				SummonerName: "上官鹰",
			},
		},
		{
			ParticipantID: 5,
			Player: lcu.Player{
				SummonerName: "封寒",
			},
		},
		{
			ParticipantID: 6,
			Player: lcu.Player{
				SummonerName: "粱萧",
			},
		},
		{
			ParticipantID: 7,
			Player: lcu.Player{
				SummonerName: "花生",
			},
		},
		{
			ParticipantID: 8,
			Player: lcu.Player{
				SummonerName: "梁思禽",
			},
		},
		{
			ParticipantID: 9,
			Player: lcu.Player{
				SummonerName: "乐之扬",
			},
		},
		{
			ParticipantID: 10,
			Player: lcu.Player{
				SummonerName: "陆渐",
			},
		},
	},
	Participants: []lcu.Participants{
		{ParticipantID: 1, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 1,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 2, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 2,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 3, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 3,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 4, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 4,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 5, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 6, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 7, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 8, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 9, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 10, Spell1ID: randSpell(), Spell2ID: randSpell(), Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
	},
}

var mockMatchList = map[string]lcu.MatchInfo{
	"10001": yiVsGe,
	"10002": yiVsGe2,
}

var yiVsGe2 = lcu.MatchInfo{
	GameDuration: 1800,
	GameID:       10002,
	ParticipantIdentities: []lcu.ParticipantIdentities{
		{
			ParticipantID: 1,
			Player: lcu.Player{
				SummonerName: "万归藏",
			},
		},
		{
			ParticipantID: 2,
			Player: lcu.Player{
				SummonerName: "朱元璋",
			},
		},
		{
			ParticipantID: 3,
			Player: lcu.Player{
				SummonerName: "云殊",
			},
		},
		{
			ParticipantID: 4,
			Player: lcu.Player{
				SummonerName: "公苏羊",
			},
		},
		{
			ParticipantID: 5,
			Player: lcu.Player{
				SummonerName: "封寒",
			},
		},
		{
			ParticipantID: 6,
			Player: lcu.Player{
				SummonerName: "粱萧",
			},
		},
		{
			ParticipantID: 7,
			Player: lcu.Player{
				SummonerName: "花生",
			},
		},
		{
			ParticipantID: 8,
			Player: lcu.Player{
				SummonerName: "梁思禽",
			},
		},
		{
			ParticipantID: 9,
			Player: lcu.Player{
				SummonerName: "乐之扬",
			},
		},
		{
			ParticipantID: 10,
			Player: lcu.Player{
				SummonerName: "陆渐",
			},
		},
	},
	Participants: []lcu.Participants{
		{ParticipantID: 1, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 1, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 1,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 2, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 2, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 2,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 3, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 3, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 3,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 4, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 4, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 4,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 5, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 5, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 6, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 6, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 7, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 7, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 8, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 8, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 9, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 9, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
		{ParticipantID: 10, Spell1ID: randSpell(), Spell2ID: randSpell(), ChampionID: 10, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5,
			Item0: randItem(), Item1: randItem(), Item2: randItem(), Item3: randItem(), Item4: randItem(), Item5: randItem(), Item6: randItem()},
		},
	},
}

func randItem() int {
	return grand.N(3050, 3090)
}

func randSpell() int {
	n := grand.N(0, 9)
	list := []int{4, 14, 11, 6, 12, 21, 3, 1, 7, 32}
	return list[n]
}
