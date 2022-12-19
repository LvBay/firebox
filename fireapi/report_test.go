package fireapi

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
)

var (
	fc = FireClient{
		Username: "test",
		Region:   "艾欧尼亚",
		Level:    "黄金",
		Domain:   "http://www.firefire.club/api",
		// Domain:   "http://127.0.0.1:8000",
	}
)

func TestCreateReport(t *testing.T) {
	r := Report{ReportedName: "摆烂大王", Reason: "firebox", ExtInfo: gconv.String(g.Map{"gameId": "12323"})}
	fc.CreateReport(r)
}
