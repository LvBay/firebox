package controller

import "github.com/LvBay/firebox/lcu"

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
		{ParticipantID: 1, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 1}},
		{ParticipantID: 2, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 2}},
		{ParticipantID: 3, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 3}},
		{ParticipantID: 4, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 4}},
		{ParticipantID: 5, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 6, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 7, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 8, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 9, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 10, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
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
		{ParticipantID: 1, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 1}},
		{ParticipantID: 2, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 2}},
		{ParticipantID: 3, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 3}},
		{ParticipantID: 4, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 4}},
		{ParticipantID: 5, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 6, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 7, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 8, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 9, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
		{ParticipantID: 10, Stats: lcu.Stats{Kills: 10, Deaths: 1, Assists: 5}},
	},
}
