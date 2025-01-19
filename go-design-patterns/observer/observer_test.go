package observer

import (
	"context"
	"testing"
	"time"
)

func Test_syncEventBus(t *testing.T) {
	observerA := NewBaseObserver("a")
	observerB := NewBaseObserver("b")
	observerC := NewBaseObserver("c")
	observerD := NewBaseObserver("d")
	bus := NewSyncEventBus()
	topic := "order_finish"
	bus.Subscribe(topic, observerA)
	bus.Subscribe(topic, observerB)
	bus.Subscribe(topic, observerC)
	bus.Subscribe(topic, observerD)

	bus.Publish(context.Background(), &Event{
		Topic: topic,
		Val:   "order_id: xxx",
	})
}

func Test_asyncEventBus(t *testing.T) {
	observerA := NewBaseObserver("a")
	observerB := NewBaseObserver("b")
	observerC := NewBaseObserver("c")
	observerD := NewBaseObserver("d")

	abus := NewAsyncEventBus()
	defer abus.Stop()

	topic := "order_finish"
	abus.Subscribe(topic, observerA)
	abus.Subscribe(topic, observerB)
	abus.Subscribe(topic, observerC)
	abus.Subscribe(topic, observerD)

	abus.Publish(context.Background(), &Event{
		Topic: topic,
		Val:   "order_id: xxx",
	})

	<-time.After(time.Second)
}
