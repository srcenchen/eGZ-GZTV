package live

import (
	"context"
	"eGZ-GZTV/api/live/v1"
	"eGZ-GZTV/internal/dao"
	"eGZ-GZTV/internal/model/entity"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

func (c *ControllerV1) GetLiveList(_ context.Context, _ *v1.GetLiveListReq) (res *v1.GetLiveListRes, err error) {
	result := dao.GetLives()
	newResult := make([]entity.LiveTable, 0)
	// 历遍所有直播，如果直播已经结束，就将直播状态改为已结束
	for _, live := range result {
		streamName := live.StreamName
		// 通过API获取直播状态
		resp, _ := http.Get("http://127.0.0.1:6022/api/stat/group?stream_name=" + streamName)
		jsonData, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		resultCode := gjson.Get(string(jsonData), "error_code")
		// 篡改直播状态
		if resultCode.Int() == 0 {
			live.LiveState = true
		}
		newResult = append(newResult, live)
	}
	res = &v1.GetLiveListRes{
		List: newResult,
	}
	return
}
