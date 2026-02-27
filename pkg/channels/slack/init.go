package slack

import (
	"github.com/andranikasd/picoclaw/pkg/bus"
	"github.com/andranikasd/picoclaw/pkg/channels"
	"github.com/andranikasd/picoclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("slack", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewSlackChannel(cfg.Channels.Slack, b)
	})
}
