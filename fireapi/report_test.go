package fireapi

import (
	"github.com/LvBay/gf/v2/frame/g"
	"github.com/LvBay/gf/v2/util/gconv"
	"testing"
)

var (
	fc = FireClient{
		Username: "test",
		Region:   "艾欧尼亚",
		Level:    "黄金",
	}
)

func TestCreateReport(t *testing.T) {
	r := Report{ReportedName: "摆烂大王", Reason: "firebox", ExtInfo: gconv.String(g.Map{"gameId": "12323"})}
	fc.CreateReport(r)
}
