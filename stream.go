package iface

import (
	"context"

	"github.com/SJTU-OpenNetwork/go-stream"
)

type StreamAPI interface {
	StartStream(context.Context, *stream.Stream) error
	AddStreamBlock(context.Context, *stream.StreamBlock) error
	SubscribeStream(context.Context, *stream.StreamConfig) error
}
