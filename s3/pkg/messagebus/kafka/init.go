package kafka

import (
	bus "github.com/opensds/crystal/s3/pkg/messagebus"
	"github.com/opensds/crystal/s3/pkg/messagebus/types"
)

func init() {
	kafkaBuilder := &KafkaBuilder{}

	bus.AddMsgBuilder(types.MSG_BUS_SENDER_KAFKA, kafkaBuilder)
}
