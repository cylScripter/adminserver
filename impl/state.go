package impl

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cylScripter/apiopen/rabbitmq"
)

type State struct {
	MqGroup *rabbitmq.MqGroup
}

var state *State

func InitState() {
	var err error
	state = &State{}
	state.MqGroup, err = rabbitmq.New()
	if err != nil {
		klog.Errorf("New rabbitmq failed %v", err)
		return
	}

	mqNode, err := rabbitmq.NewRabbitMQNode(rabbitmq.NodeConfig{
		Host:     "localhost",
		Password: "guest",
		Port:     5672,
		User:     "guest",
		VHost:    "",
	})
	if err != nil {
		klog.Errorf("New rabbitmq node failed %v", err)
		return
	}
	state.MqGroup.AddNode(mqNode)

}

func GetState() *State {
	return state
}
