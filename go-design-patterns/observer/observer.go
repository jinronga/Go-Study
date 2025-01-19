package observer

import (
	"context"
	"fmt"
	"sync"
)

// 观察者模式

type (
	Event struct {
		Topic string
		Val   interface{}
	}

	Observer interface {
		OnChange(ctx context.Context, e *Event) error
	}

	EventBus interface {
		Subscribe(topic string, o Observer)
		Unsubscribe(topic string, o Observer)
		Publish(ctx context.Context, e *Event)
	}

	BaseObserver struct {
		name string
	}

	BaseEventBus struct {
		mux       sync.RWMutex
		observers map[string]map[Observer]struct{}
	}

	SyncEventBus struct {
		BaseEventBus
	}

	observerWithErr struct {
		o   Observer
		err error
	}

	AsyncEventBus struct {
		BaseEventBus
		errC chan *observerWithErr
		ctx  context.Context
		stop context.CancelFunc
	}
)

func NewBaseObserver(name string) *BaseObserver {
	return &BaseObserver{
		name: name,
	}
}
func (b *BaseObserver) OnChange(ctx context.Context, e *Event) error {
	fmt.Printf("observer: %s, event key: %s, event val: %v", b.name, e.Topic, e.Val)
	fmt.Println("=======================")
	return nil
}

func NewBaseEventBus() BaseEventBus {
	return BaseEventBus{
		observers: make(map[string]map[Observer]struct{}),
	}
}

func (b *BaseEventBus) Subscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	_, ok := b.observers[topic]
	if !ok {
		b.observers[topic] = make(map[Observer]struct{})
	}
	b.observers[topic][o] = struct{}{}
}

func (b *BaseEventBus) Unsubscribe(topic string, o Observer) {
	b.mux.Lock()
	defer b.mux.Unlock()
	delete(b.observers[topic], o)
}

func NewSyncEventBus() *SyncEventBus {
	return &SyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
}

func (s *SyncEventBus) Publish(ctx context.Context, e *Event) {
	s.mux.RLock()
	subscribers := s.observers[e.Topic]
	s.mux.RUnlock()

	errs := make(map[Observer]error)
	for subscriber := range subscribers {
		if err := subscriber.OnChange(ctx, e); err != nil {
			errs[subscriber] = err
		}
	}

	s.handleErr(ctx, errs)
}

func (s *SyncEventBus) handleErr(_ context.Context, errs map[Observer]error) {
	for o, err := range errs {
		// 处理 publish 失败的 observer
		fmt.Printf("observer: %v, err: %v", o, err)
	}
}

func NewAsyncEventBus() *AsyncEventBus {
	aBus := AsyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
	aBus.ctx, aBus.stop = context.WithCancel(context.Background())
	// 处理处理错误的异步守护协程
	go aBus.handleErr()
	return &aBus
}

func (a *AsyncEventBus) Stop() {
	a.stop()
}

func (a *AsyncEventBus) Publish(ctx context.Context, e *Event) {
	a.mux.RLock()
	subscribers := a.observers[e.Topic]
	defer a.mux.RUnlock()
	for sub := range subscribers {
		// shadow
		subscriber := sub
		go func() {
			if err := subscriber.OnChange(ctx, e); err != nil {
				select {
				case <-a.ctx.Done():
				case a.errC <- &observerWithErr{
					o:   subscriber,
					err: err,
				}:
				}
			}
		}()
	}
}

func (a *AsyncEventBus) handleErr() {
	for {
		select {
		case <-a.ctx.Done():
			return
		case resp := <-a.errC:
			// 处理 publish 失败的 observer
			fmt.Printf("observer: %v, err: %v", resp.o, resp.err)
		}
	}
}
