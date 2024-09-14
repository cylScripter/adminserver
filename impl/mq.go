package impl

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cylScripter/apiopen/admin"
	"github.com/cylScripter/apiopen/rabbitmq"
)

var TestMq *rabbitmq.JsonMsg

func InitMq() {
	TestMq = rabbitmq.NewJsonMsg(GetState().MqGroup, "", &admin.GetUserListReq{})

	mq := rabbitmq.NewConsumer(GetState().MqGroup, 0)

	err := mq.AddQueue(nil, rabbitmq.Queue{
		Name:        "getUserList",
		HandlerFunc: TestCMq,
	})
	if err != nil {
		klog.Errorf("Add queue failed %v", err)
	}

}

func TestCMq(ctx *context.Context, req *rabbitmq.ConsumeReq) error {
	klog.Infof("consume msg %s", req.Data)
	klog.Infof("consume msg %v", req.MsgId)
	klog.Infof("consume msg %v", req.CreatedAt)

	return nil
}
