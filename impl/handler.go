package impl

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	admin "github.com/cylScripter/apiopen/admin"
	"github.com/cylScripter/apiopen/rabbitmq"
)

// AdminImpl implements the last service interface defined in the IDL.
type AdminImpl struct{}

// GetUserList implements the AdminImpl interface.
func (s *AdminImpl) GetUserList(ctx context.Context, req *admin.GetUserListReq) (*admin.GetUserListResp, error) {
	var resp admin.GetUserListResp
	err := TestMq.Pub(&ctx, &rabbitmq.PubReq{
		Group:     0,
		RouterKey: "getUserList",
	}, req)
	if err != nil {
		klog.Errorf("publish error %v", err)
		return nil, err
	}

	return &resp, nil
}
