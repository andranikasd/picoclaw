package telegram

import (
	"github.com/andranikasd/picoclaw/pkg/bus"
	"github.com/andranikasd/picoclaw/pkg/channels"
	"github.com/andranikasd/picoclaw/pkg/config"
)

func init() {
	channels.RegisterFactory("telegram", func(cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
		return NewTelegramChannel(cfg, b)
	})
}
