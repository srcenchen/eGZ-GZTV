package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"eGZ-GZTV/api/user/v1"
)

func (c *ControllerV1) LoginVerify(ctx context.Context, req *v1.LoginVerifyReq) (res *v1.LoginVerifyRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
