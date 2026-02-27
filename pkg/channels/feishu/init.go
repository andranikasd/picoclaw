package feishu

import (
	"github.com/andranikasd/picoclaw/pkg/bus"
	"github.com/andranikasd/picoclaw/pkg/channels"
	"github.com/andranikasd/picoclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("feishu", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewFeishuChannel(cfg.Channels.Feishu, b)
	})
}
