package live

import (
	"context"
	"eGZ-GZTV/internal/dao"
	"github.com/tidwall/gjson"
	"io"
	"net/http"

	"eGZ-GZTV/api/live/v1"
)

func (c *ControllerV1) GetLive(_ context.Context, req *v1.GetLiveReq) (res *v1.GetLiveRes, err error) {
	result := dao.GetLiveByID(req.Id)
	// 历遍所有直播，如果直播已经结束，就将直播状态改为已结束
	streamName := result.StreamName
	// 通过API获取直播状态
	resp, _ := http.Get("http://127.0.0.1:6022/api/stat/group?stream_name=" + streamName)
	jsonData, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	resultCode := gjson.Get(string(jsonData), "error_code")
	// 篡改直播状态
	if resultCode.Int() == 0 {
		result.LiveState = true
	}

	res = &v1.GetLiveRes{
		Live: result,
	}
	return
}
