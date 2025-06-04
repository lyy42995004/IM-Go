package kafka

import (
	"strings"

	"github.com/IBM/sarama"
	"github.com/lyy42995004/IM-Go/pkg/log"
)

var producer sarama.AsyncProducer
var topic string = "default_message"

func InitProducer(topicInput, hosts string) {
	topic = topicInput
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP

	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		log.Error("init kafka client error", log.Err(err))
	}

	producer, err = sarama.NewAsyncProducerFromClient(client)
	if nil != err {
		log.Error("init kafka async client error", log.Err(err));
	}
}

func Send(data []byte) {
	be := sarama.ByteEncoder(data)
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: be}
}

func Close() {
	if producer != nil {
		producer.Close()
	}
}
