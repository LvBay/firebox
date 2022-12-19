package controller

import (
	"context"
	"github.com/LvBay/firebox/lcu"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMatchInfoReq struct {
	g.Meta `path:"/lol-match-history/v1/games/{id}" method:"get" summary:"根据gameid获取对局信息"`
	Id     string
}

type GetMatchInfoRes struct {
	lcu.MatchInfo
}

type GetCurrMatchListReq struct {
	g.Meta `path:"/lol-match-history/v3/matchlist/account/{id}" method:"get" summary:"根据puuid获取对局记录"`
	Id     string
}

type GetCurrMatchListRes struct {
	lcu.MatchList
}

func (s *service) GetMatchInfo(ctx context.Context, req *GetMatchInfoReq) (*GetMatchInfoRes, error) {
	m, ok := mockMatchList[req.Id]
	if ok {
		return &GetMatchInfoRes{m}, nil
	}
	ret := &GetMatchInfoRes{yiVsGe}
	return ret, nil
}

func (s *service) GetCurrMatchList(ctx context.Context, req *GetCurrMatchListReq) (*GetCurrMatchListRes, error) {
	ret := &GetCurrMatchListRes{
		lcu.MatchList{
			Games: lcu.Games{
				Games: []lcu.MatchInfo{
					yiVsGe,
					yiVsGe2,
				},
			},
		},
	}
	return ret, nil
}
