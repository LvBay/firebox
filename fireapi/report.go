package fireapi

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"log"
)

type Report struct {
	Region       string `json:"region"        description:"大区##"`        // 大区##
	Reason       string `json:"reason"        description:"理由##"`        // 理由##
	Level        string `json:"level"        description:"段位##"`         // 理由##
	ReportType   int    `json:"reportType"   description:"类型，1举报，2表扬##"` // 类型，1举报，2表扬##
	ExtraGiftId  int    `json:"extraGiftId" description:"额外惩罚/奖励##"`     // 额外惩罚/奖励##
	ReportedName string `json:"reportedName" description:"被通报游戏玩家名称##"`  // 被通报游戏玩家名称##
	CreatorName  string `json:"creatorName"  description:"创建人的游戏id##"`   // 创建人的游戏id##
	ExtInfo      string `json:"extInfo"      description:""`             //
}

var (
	ctx    = gctx.New()
	client = gclient.New()
)

func (c *FireClient) CreateReport(report Report) error {
	report.CreatorName = c.Username
	report.Region = c.Region
	report.Level = c.Level
	u := fmt.Sprintf("%s/v2/report", c.Domain)
	resp := client.PostVar(context.Background(), u, report)
	log.Println(resp.String())

	return nil
}

type FireClient struct {
	Domain   string
	Username string
	Region   string
	Level    string
}
