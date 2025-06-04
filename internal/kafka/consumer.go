package kafka

import (
	"strings"

	"github.com/IBM/sarama"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

var consumer sarama.Consumer

type ConsumerCallBack func(data []byte)

func InitConsumer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, "."), config)
	if err != nil {
		log.Error("init kafka consumer error", log.Err(err))
	}

	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Error("init kafka consumer error", log.Err(err))
	}
}

// 消费消息
func ConsumerMsg(cb ConsumerCallBack) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Error("ConsumePartition error", log.Err(err))
	}

	defer partitionConsumer.Close()
	for {
		msg := <-partitionConsumer.Messages()
		if cb != nil {
			cb(msg.Value)
		}
	}
}

func CloseConsumer() {
	if nil != consumer {
		consumer.Close()
	}
}
