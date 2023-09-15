package messagebus

import (
	"github.com/opensds/crsytal/s3/pkg/helper"
)

type MessageSenderBuilder interface {
	Create(config helper.MsgBusConfig) (MessageSender, error)
}

var MsgBuilders = make(map[int]MessageSenderBuilder)

func AddMsgBuilder(builderType int, builder MessageSenderBuilder) {
	MsgBuilders[builderType] = builder
}
