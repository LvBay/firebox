package controller

import (
	"context"
	"github.com/LvBay/firebox/lcu"
	"github.com/LvBay/gf/v2/frame/g"
)

type CurrentSummonerReq struct {
	g.Meta `path:"/lol-summoner/v1/current-summoner" method:"get" summary:"获取当前召唤师"`
}

type CurrentSummonerRes struct {
	lcu.Summoner
}

var Service = service{}

type service struct{}

func (s *service) GetCurrentSummoner(ctx context.Context, req *CurrentSummonerReq) (*CurrentSummonerRes, error) {
	res := &CurrentSummonerRes{lcu.Summoner{
		DisplayName:   "测试呀",
		SummonerLevel: 100,
	}}
	return res, nil
}
