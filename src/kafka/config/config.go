package config

const (
	ProducerPort       = ":9812"
	KafkaServerAddress = "confluent-kafka-tw-1.int.onemount.dev:9093,confluent-kafka-tw-2.int.onemount.dev:9093,confluent-kafka-tw-3.int.onemount.dev:9093,confluent-kafka-sg-4.int.onemount.dev:9093,confluent-kafka-sg-5.int.onemount.dev:9093,confluent-kafka-sg-6.int.onemount.dev:9093"
	KafkaTopic         = "tcb-promotion-rule-common-job-qc"
)

const (
	ConsumerGroup = "tcb-promotion-rule-common-job-qc"
	ConsumerTopic = "tcb-promotion-rule-common-job-qc"
	ConsumerPort  = ":8081"
)
