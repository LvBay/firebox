package controller

import (
	"context"
	"github.com/LvBay/firebox/lcu"
	"github.com/gogf/gf/v2/frame/g"
)

type GetGameFlowSessionReq struct {
	g.Meta `path:"/lol-gameflow/v1/session" method:"get" summary:"获取双方召唤师信息"`
}

type GetGameFlowSessionRes struct {
	lcu.GameFlowSession
}

func (s *service) GetGameFlowSession(ctx context.Context, req *GetGameFlowSessionReq) (*GetGameFlowSessionRes, error) {
	return &GetGameFlowSessionRes{
		lcu.GameFlowSession{
			GameData: lcu.GameData{
				PlayerChampionSelections: []lcu.PlayerChampionSelections{
					{SummonerInternalName: "Aphelios", ChampionId: 523},
					{SummonerInternalName: "Rell", ChampionId: 526},
					{SummonerInternalName: "Pyke", ChampionId: 555},
					{SummonerInternalName: "Vex", ChampionId: 711},
					{SummonerInternalName: "Yone", ChampionId: 777},
					{SummonerInternalName: "Sett", ChampionId: 875},
					{SummonerInternalName: "Lillia", ChampionId: 876},
					{SummonerInternalName: "Gwen", ChampionId: 887},
					{SummonerInternalName: "Renata", ChampionId: 888},
					{SummonerInternalName: "Nilah", ChampionId: 895},
				},
				TeamOne: []lcu.Team{
					{AccountId: 11, SummonerName: "梁思11", SummonerInternalName: "Aphelios", SelectedPosition: "TOP"},
					{AccountId: 12, SummonerName: "梁思12", SummonerInternalName: "Rell", SelectedPosition: "JUNGLE"},
					{AccountId: 14, SummonerName: "梁思14", SummonerInternalName: "Vex", SelectedPosition: "BOTTOM"},
					{AccountId: 15, SummonerName: "梁思15", SummonerInternalName: "Yone", SelectedPosition: "UTILITY"},
					{AccountId: 13, SummonerName: "梁思13", SummonerInternalName: "Pyke", SelectedPosition: "MIDDLE"},
				},
				TeamTwo: []lcu.Team{
					{AccountId: 21, SummonerName: "梁思21", SummonerInternalName: "Sett", SelectedPosition: "TOP"},
					{AccountId: 22, SummonerName: "梁思22", SummonerInternalName: "Lillia", SelectedPosition: "JUNGLE"},
					{AccountId: 24, SummonerName: "梁思24", SummonerInternalName: "Renata", SelectedPosition: "BOTTOM"},
					{AccountId: 25, SummonerName: "梁思25", SummonerInternalName: "Nilah", SelectedPosition: "UTILITY"},
					{AccountId: 23, SummonerName: "梁思23", SummonerInternalName: "Gwen", SelectedPosition: "MIDDLE"},
				},
			},
		},
	}, nil
}
