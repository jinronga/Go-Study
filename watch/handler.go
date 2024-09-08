package watch

import (
	"context"
	"sync"
)

type (

	// Handler 消息处理
	Handler interface {
		// Handle 处理消息
		//
		// 	ctx 上下文
		// 	msg 消息
		Handle(ctx context.Context, msg *Message) error
	}

	// HandleFun 消息处理函数
	HandleFun func(ctx context.Context, msg *Message) error

	// defaultHandler 默认消息处理
	defaultHandler struct {
		lock           sync.Mutex
		topicHandleMap map[vobj.Topic][]HandleFun
	}

	// DefaultHandlerOption 默认消息处理配置
	DefaultHandlerOption func(d *defaultHandler)
)
