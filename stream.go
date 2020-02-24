package iface

import (
	"context"
    peer "github.com/libp2p/go-libp2p-core/peer"
	"github.com/SJTU-OpenNetwork/go-stream"
)

type StreamAPI interface {
	StartStream(context.Context, *stream.Stream) error
	AddStreamBlock(context.Context, *stream.StreamBlock) error
	SubscribeStream(context.Context, *stream.StreamConfig) error
    StartWorker(context.Context, *stream.StreamConfig, peer.ID) error
}
