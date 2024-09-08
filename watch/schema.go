package watch

type (
	// Schemer 消息编解码器
	Schemer interface {
		// Decode 解码
		Decode(in *Message, out any) error

		// Encode 编码
		Encode(in *Message, out any) error
	}

	defaultSchemer struct{}

	emptySchemer struct{}
)
