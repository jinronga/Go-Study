package vobj

// Topic 消息类型
//
//go:generate
type Topic int

const (
	// TopicUnknown 未知
	TopicUnknown Topic = iota // 未知

	// TopicStrategy 策略
	TopicStrategy // 策略

	// TopicAlert 单条告警
	TopicAlert // 单条告警

	// TopicAlarm 多条告警
	TopicAlarm // 多条告警
)
